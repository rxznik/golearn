# Weather CLI

**Приложение командной строки, позволяющее узнать погоду на сегодня в любом городе мира.**

### ⬇️ Установка и запуск ⬇️

<br/>

<details>
<summary><strong>⚙️ By Git + Go</strong></summary>

<br/>

*Клонируем репозиторий*:

```bash
git clone https://github.com/rxznik/golearn.git
```

*Переходим в директорию проекта*:

```bash
cd golearn/weather-cli
```

*Устанавливаем зависимости*:

```bash
go mod download
```

*Запускаем приложение*:

```bash
go run cmd/main.go --help
```

*Пример получения погоды в Москве с включенным логированием*:

```bash
go run cmd/main.go -l true Москва
```

</details>

<br/>

<details>
<summary><strong>🧰 By Git + Task</strong></summary>

<br/>

*Клонируем репозиторий*:

```bash
git clone https://github.com/rxznik/golearn.git
```

*Переходим в директорию проекта*:

```bash
cd golearn/weather-cli
```

*Устанавливаем зависимости*:

```bash
task download
```

*Собираем приложение*:

```bash
task build
```

*Запускаем приложение*:

```bash
task run -- --help
```

*Пример получения погоды в Москве с включенным логированием*:

```bash
task run -- -l true Москва
```

</details>

<br/>

<details>
<summary><strong>📦 By Git + Docker</strong></summary>

<br/>

*Клонируем репозиторий*:

```bash
git clone https://github.com/rxznik/golearn.git
```

*Переходим в директорию проекта*:

```bash
cd golearn/weather-cli
```

*Устанавливаем зависимости*:

```bash
docker build -t weather-cli:latest -f ./build/Dockerfile .
```

*Запускаем приложение*:

```bash
docker run --rm weather-cli:latest
```

*Пример получения погоды в Москве с включенным логированием*:

```bash
docker run --rm weather-cli:latest -l true Москва
```

</details>

<br/>

<details>
<summary><strong>🐋 By pull from Docker Hub</strong></summary>

<br/>

*Скачиваем образ из Docker Hub*:

```bash
docker pull rxznik/weather-cli:latest
```

*Запускаем приложение*:

```bash
docker run --rm rxznik/weather-cli:latest
```

*Пример получения погоды в Москве с включенным логированием*:

```bash
docker run --rm rxznik/weather-cli:latest -l true Москва
```

</details>

### Использование

*Получить справочную информацию*:

```bash
# через go
go run cmd/main.go --help

# через docker
docker run --rm rxznik/weather-cli:latest

# второй вариант через docker
docker run --rm rxznik/weather-cli:latest --help
```

*Пример получения погоды в Москве*:

```bash
# через go
go run cmd/main.go Москва

# через docker
docker run --rm rxznik/weather-cli:latest Москва
```

*Пример получения погоды в Москве с включенным логированием*:

```bash
# через go
go run cmd/main.go -l true Москва

# или с полным названием флага
go run cmd/main.go --loud true Москва

# через docker
docker run --rm rxznik/weather-cli:latest -l true Москва
```

### 📜 API и пакеты

* [OpenMeteo API](https://api.open-meteo.com/)

* [urfave/cli/v2](https://github.com/urfave/cli)

* [zap](https://pkg.go.dev/go.uber.org/zap)

* [cleanenv](https://github.com/ilyakaznacheev/cleanenv)

* [testify](https://pkg.go.dev/github.com/stretchr/testify)