# Конкурентность и каналы в Go

<details>
<summary><strong>🏠 Конкурентность и параллелизм</strong></summary>

### Определения
- **Конкурентность**: Логическая возможность выполнять задачи _независимо_, даже если они выполняются на одном ядре CPU.
- **Параллелизм**: Физическое выполнение задач _одновременно_ на нескольких ядрах CPU.

### Реализация в Go
- **Горутины**: Легковесные потоки (2 КБ стека, динамически расширяются).
- **Переключение горутин**:
  - **Системные вызовы** (блокирующие I/O, sleep).
  - **Блокировка каналов** (отправка/прием).
  - **Явный вызов** `runtime.Gosched()`.
  - **Preemption** от Sysmon (при долгих вычислениях >10 мс).

```go
// Пример переключения через sysmon
func infiniteLoop() {
    for { /* Sysmon прервет через 10 мс */ }
}
```

### Когда срабатывает переключение?
- **Кооперативная многозадачность**: Горутина сама отдает управление (каналы, syscall).
- **Вытесняющая**: Sysmon принудительно останавливает "жадные" горутины.
</details>

<details>
<summary><strong>⚙️ Планировщик Go (Scheduler)</strong></summary>

### Модель PMG
- **P (Processor)**: Логический процессор (количество = `GOMAXPROCS`).
- **M (Machine)**: Поток ОС (управляется планировщиком).
- **G (Goroutine)**: Задача.

### Алгоритм работы P
1. **Локальная очередь (LRQ)**: Обрабатывает свои G.
2. **Глобальная очередь (GRQ)**: Берет 1 G каждые 61 шаг (для балансировки).
3. **Work-stealing**: Крадет 50% G из LRQ другого P.

```go
// Пример работы P
runtime.GOMAXPROCS(4) // 4 P создадут 5 системных потоков
```

### Механизмы
- **Handoff**: При блокирующем syscall P освобождает M, чтобы другие G могли использовать ядра.
- **Sysmon**: Мониторит:
    - Сетевые операции через **Netpoller** (epoll/kqueue).
    - Долгие G (>10 мс) → preemption.
- **Netpoller**: Асинхронно обрабатывает сетевые вызовы, не блокируя M.

### Очереди
- **LRQ**: Локальная очередь P (до 256 G).
- **GRQ**: Глобальная очередь для новых G.

### Потоки (M)
- Стартуют по необходимости (до лимита `ulimit -n`).
- При блокировке (syscall) → новый M создается для других G.

### Как горутины попадают в Netpoller?
1. **Сетевой вызов**: Горутина выполняет операцию (например, `conn.Read()`).
2. **Неблокирующий режим**: Go автоматически переводит сокет в неблокирующий режим.
3. **Регистрация в Netpoller**:
  - Файловый дескриптор сокета добавляется в `epoll/kqueue/IOCP`.
  - Горутина переводится в состояние **ожидания** (Gwaiting).
4. **Освобождение ресурсов**:
  - P открепляется от M (если это системный вызов).
  - M может выполнять другие горутины.

### Как горутины выходят из Netpoller?
1. **Событие готовности**: ОС уведомляет Netpoller о готовности сокета (данные пришли, можно писать).
2. **Пробуждение горутины**:
  - Netpoller помечает горутину как **готовую к выполнению**.
  - Горутина добавляется в **глобальную очередь (GRQ)** P.
3. **Планирование**:
  - Когда P получит управление, он начнет выполнять пробужденную горутину.

### Какой P получает пробужденную горутину?
Когда Netpoller (через `epoll/kqueue/IOCP`) обнаруживает готовность сокета, пробужденная горутина **не привязана жестко к конкретному P**.
Алгоритм распределения зависит от контекста:

1. **Общий случай**:
  - Горутина помещается в **глобальную очередь (GRQ)**.
  - Любой свободный P может забрать её через механизм **work-stealing** или при обработке GRQ (каждые 61 шаг).

2. **Оптимизация для привязки к исходному P**:
  - Если горутина была заблокирована **на короткое время** (например, быстрое сетевое событие),
    runtime пытается вернуть её в **локальную очередь (LRQ) исходного P** (если он активен).
  - Это улучшает локальность данных и снижает накладные расходы.

3. **Системный монитор (Sysmon)**:
  - Sysmon периодически проверяет GRQ и **равномерно распределяет горутины** по LRQ свободных P.
  - Это предотвращает "голодание" отдельных P.

```go
// Пример: Чтение из сети с неявным использованием Netpoller
conn, _ := net.Dial("tcp", "example.com:80")
buf := make([]byte, 1024)
n, _ := conn.Read(buf) // Блокировка только горутины, не потока!
```

**Асинхронная обработка**: Все сетевые вызовы в Go по умолчанию не блокируют потоки ОС.

#### Детали реализации
 - При вызове `conn.Read()` **runtime** вызывает `runtime.netpollblock()`.

 - Данные о горутине сохраняются в структуре сокета.

 - После события **epoll_wait** горутина помечается как **Runnable**.
</details>

<details>
<summary><strong>🌉 Каналы (Channels)</strong></summary>

### Структура (runtime.hchan)
```go
type hchan struct {
	qcount   uint           // кол-во данных в кольцевой очереди
	// other fields (dataqsize, elemtype, elemsize)
	closed   uint32         // закрыт ли канал, uint32 из-за атомарных операций
    buf      unsafe.Pointer // кольцевой буфер
    sendx    uint           // индекс отправки
    recvx    uint           // индекс приема
    lock     mutex          // мьютекс
    sendq    waitq          // очередь ожидающих отправителей
    recvq    waitq          // очередь ожидающих получателей
}
```

### Небуферизированный канал
- **Отправка**: Блокирует отправителя, пока получатель не готов.
- **Прием**: Блокирует получателя, пока отправитель не готов.

```go
ch := make(chan int)
go func() { ch <- 1 }() // Блокируется, пока main не прочитает
fmt.Println(<-ch)       // Разблокирует отправителя
```

### Буферизированный канал
- **Отправка**: Не блокирует, пока буфер не заполнен.
- **Прием**: Не блокирует, пока буфер не пуст.

```go
ch := make(chan int, 2)
ch <- 1  // OK
ch <- 2  // OK
ch <- 3  // Блокировка (буфер заполнен)
```

### Select
- Обрабатывает первый готовый канал.
- **Non-blocking** с `default`.

```go
select {
case v := <-ch: // Чтение
case ch <- 10:   // Запись
default:         // Неблокирующий режим
}
```

### Под капотом
- **Блокировка**: Горутина попадает в `sendq` или `recvq`.
- **Пробуждение**: При появлении пары (отправитель/получатель).
</details>

<details>
<summary><strong>🧰 Примитивы синхронизации</strong></summary>

### WaitGroup
- **Цель**: Ожидание завершения группы горутин.
- **Плюсы**: Простота использования.
- **Минусы**: Нельзя переиспользовать без `Add`.

```go
var wg sync.WaitGroup
wg.Add(2)
go func() { defer wg.Done() }()
go func() { defer wg.Done() }()
wg.Wait()
```

### Mutex
- **Цель**: Исключение гонок данных.
- **Плюсы**: Точный контроль.
- **Минусы**: Риск дедлоков.

```go
var mu sync.Mutex
mu.Lock()
counter++
mu.Unlock()
```

### RWMutex
- **Чтение**: Множественный доступ.
- **Запись**: Эксклюзивный доступ.

### Atomic
- **Цель**: Атомарные операции без блокировок.
- **Плюсы**: Высокая скорость.
- **Минусы**: Только для примитивов (int32, pointers).

```go
var counter int32
atomic.AddInt32(&counter, 1)
```
</details>

<details>
<summary><strong>🎯 Механизмы для работы с конкурентностью</strong></summary>

### SingleFlight (golang.org/x/sync/singleflight)
- **Цель**: Предотвращение повторных вычислений для одинаковых запросов.
- **Принцип**: Группировка одновременных вызовов с одним ключом.

```go
var group singleflight.Group
result, _ := group.Do("key", func() (interface{}, error) {
    return fetchFromDB()
})
```
[Реализация](https://github.com/rxznik/golearn/blob/main/concurrency/balun/channels/singleflight/main.go)

### RateLimiter (golang.org/x/time/rate)
- **Цель**: Ограничение частоты запросов (например, API).
- **Принцип**: Токенный алгоритм (token bucket).

```go
limiter := rate.NewLimiter(rate.Every(100*time.Millisecond), 5)
if limiter.Allow() { /* Выполнить */ }
```
[Реализация](https://github.com/rxznik/golearn/blob/main/concurrency/balun/channels/rate_limiter/main.go)

### ErrGroup (golang.org/x/sync/errgroup)
- **Цель**: Группа горутин с обработкой ошибок.
- **Принцип**: Отмена всех задач при первой ошибке.

```go
g, ctx := errgroup.WithContext(ctx)
g.Go(func() error { return nil })
if err := g.Wait(); err != nil {}
```
[Реализация](https://github.com/rxznik/golearn/blob/main/concurrency/balun/channels/errgroup/main.go)

### Семафор (реализация через каналы)
- **Цель**: Ограничение одновременных операций.

```go
sem := make(chan struct{}, 3)
sem <- struct{}{} // Захват
<-sem             // Освобождение
```
[Реализация](https://github.com/rxznik/golearn/blob/main/concurrency/balun/channels/semaphore/main.go)

### WorkerPool
- **Цель**: Пул воркеров для обработки задач.
- **Принцип**: Фиксированное количество горутин + канал задач.

```go
jobs := make(chan Task, 100)
for i := 0; i < 10; i++ {
    go func() { for task := range jobs { process(task) } }()
}
```
[Реализация](https://github.com/rxznik/golearn/blob/main/concurrency/habr/workerpool/main.go)
</details>

<details>
<summary><strong>💫 Context</strong></summary>

### Цель
- Отмена операций (например, HTTP-запросов).
- Передача данных (request-scoped данные).

### Принцип
- **Дерево контекстов**: Родительский контекст может отменить все дочерние.
- **Методы**:
    - `WithCancel` → `cancel()`.
    - `WithTimeout` → автоотмена через время.
    - `WithValue` → передача значений.

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

go func() {
    select {
    case <-ctx.Done():
        return // Прервать операцию
    }
}()
```
</details>

<details>
<summary><strong>⚒️ Паттерны работы с каналами</strong></summary>

### Fan-out
- **Цель**: Распределение задач между несколькими воркерами.
- **Реализация**: Один входной канал → N горутин.

```go
func fanOut(input <-chan int, workers int) {
    for i := 0; i < workers; i++ {
        go func() { for v := range input { process(v) } }()
    }
}
```
[Реализация](https://github.com/rxznik/golearn/blob/main/concurrency/balun/channels/fan-out/main.go)

### Fan-in
- **Цель**: Объединение результатов из нескольких каналов.
- **Реализация**: N каналов → один выходной.

```go
func fanIn(inputs ...<-chan int) <-chan int {
    out := make(chan int)
    for _, in := range inputs {
        go func(ch <-chan int) { for v := range ch { out <- v } }(in)
    }
    return out
}
```
[Реализация](https://github.com/rxznik/golearn/blob/main/concurrency/balun/channels/fan-in/main.go)

### Tee
- **Цель**: Разделение данных на два канала.

```go
func tee(input <-chan int) (_, _ <-chan int) {
    out1, out2 := make(chan int), make(chan int)
    go func() {
        for v := range input {
            out1 <- v
            out2 <- v
        }
    }()
    return out1, out2
}
```
[Реализация](https://github.com/rxznik/golearn/blob/main/concurrency/balun/channels/tee/main.go)
</details>
