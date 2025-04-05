package main

import "fmt"

type Vehicle interface {
	Move()
	Info()
	Stop()
}

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

func (c Car) Info() {
	fmt.Printf("Car %s has %d places\n", c.Name, c.Places)
}

func (c Car) Stop() {
	fmt.Printf("Car %s stopped\n", c.Name)
}

func (b Boat) Move() {
	fmt.Printf("Boat %s is moving with speed %d\n", b.Name, b.Speed)
}

func (b Boat) Info() {
	fmt.Printf("Boat %s has size %d x %d\n", b.Name, b.SizeX, b.SizeY)
}

func (b Boat) Stop() {
	fmt.Printf("Boat %s stopped\n", b.Name)
}

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
