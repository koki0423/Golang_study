package main

import "fmt"

func main() {
	num := 123
	str := "ABC"
	const pi = 3.14159

	fmt.Print("num=", num, " str=", str, "\n")
	fmt.Println("num =", num, "str =", str)
	fmt.Printf("num=%d str=%s\n", num, str)
	fmt.Println("pi =", pi)

	//pi=3 のようにconstは変更できない

	/*var [変数名] [型名] <=[数値]>*/
	var num1 int = 10
	var (
		name string = "tanaka"
		age  int    = 25
	)
	fmt.Println("num1+5=", num1+5)
	fmt.Println("name=", name, "age=", age)
}
