package main

import "fmt"

type Vehicle interface {
	Drive(km float64)
	AddFuel(litters float64)
	Efficiency() float64
}

type Car struct {
	km   float64
	fuel float64
}

type Truck struct {
	km   float64
	fuel float64
}

func (c *Car) Drive(km float64) {
	c.km += km
}

func (c *Car) AddFuel(liters float64) {
	c.fuel += liters
}

func (c *Car) Efficiency() float64 {
	if c.fuel == 0 {
		return 0 // 燃料がない場合は効率は0として返す
	}
	return c.km / c.fuel
}

func (t *Truck) Drive(km float64) {
	t.km += km
}
func (t *Truck) AddFuel(liters float64) {
	t.fuel += liters
}
func (t *Truck) Efficiency() float64 {
	if t.fuel == 0 {
		return 0
	}
	return t.km / t.fuel
}

func main() {
	c := &Car{}
	c.AddFuel(500)
	c.Drive(300)

	t := &Truck{}
	t.AddFuel(1000)
	t.Drive(600)

	printEfficiency(c)
	printEfficiency(t)
}

func printEfficiency(v Vehicle) {
	fmt.Printf("Efficiency: %f\n", v.Efficiency())
}
