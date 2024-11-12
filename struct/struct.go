package main

import "fmt"

type students struct{
	name string
	age int
	gender string
	garade int
}

func main(){
	var student students
	student.name = "John"
	student.age =21
	student.gender = "Female"
	student.garade =2

	fmt.Println(student)
}