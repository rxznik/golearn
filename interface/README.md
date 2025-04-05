# Интерфейсы в Go (interface)

<details>
<summary><b>🏠 Основные понятия</b></summary>

**Интерфейс** — контракт, которому должны соответствовать другие объекты (в случае Go — структуры), чтобы они могли взаимодействовать с ним. Интерфесы в Go помогают реализовать принцип инверсии зависимостей (Dependency Inversion principle).

Интерфейсы представляют абстракцию поведения других типов. Интерфейсы позволяют определять функции, которые не привязаны к конкретной реализации. То есть интерфейсы определяют некоторый функционал, но не реализуют его.

*Создание интерфейса:*
```go
type MyInterface interface {
    // Определение методов интерфейса
    MustMethod[T any](args ...T) T
    Method[T any](args ...T) (T, error)
    // ...
}
```

</details>

<details>
<summary><b>🎯 Примеры реализации интерфейсов в Go</b></summary>

```go
package main

import "fmt"

// Определяем интерфейс
type Vehicle interface {
	Move()
	Info()
	Stop()
}

// Определяем структуры, которые будут реализовывать интерфейс
type Car struct {
	Name   string
	Speed  int
	Places int
}

type Boat struct {
	Name  string
	Speed int
	SizeX int
	SizeY int
}

func (c Car) Move() {
	fmt.Printf("Car %s is moving with speed %d\n", c.Name, c.Speed)
}

// По аналогии реализуем остальные методы интерфейса...
// code...

func (b Boat) Stop() {
	fmt.Printf("Boat %s stopped\n", b.Name)
}

// Пример использования
func main() {
	var car Vehicle = Car{Name: "BMW", Speed: 100, Places: 4}
	var boat Vehicle = Boat{Name: "Yacht", Speed: 10, SizeX: 10, SizeY: 10}

	vehicles := []Vehicle{car, boat}
	for _, vehicle := range vehicles {
		vehicle.Info()
		drive(vehicle)
	}
}

func drive(vehicle Vehicle) {
	vehicle.Move()
	vehicle.Stop()
}
```
</details>