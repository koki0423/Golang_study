package main

import "fmt"

func add (a int, b int) int {
	return a + b
}

func main(){
	var (
		a int = 10
		b int = 20
	)
	
	fmt.Println(add(a,b))
}