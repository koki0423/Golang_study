package main

import "fmt"

type Car struct{
	make string
	model string
}

func (c Car) display() string{
	return c.make + " " + c.model
}

func main(){
	car:= new(Car)
	car.make="Toyota"
	car.model="AQUA"
	fmt.Println(car.display())
}