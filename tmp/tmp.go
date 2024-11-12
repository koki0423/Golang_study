package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scan(&a)
	slice := make([]int, 0, a*4)

	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		slice = append(slice, b)
	}

	for i := len(slice)-1; i >= 0; i-- {
		if i == 0 {
			fmt.Printf("%d\n",slice[i])
			return
		}
		fmt.Printf("%d ",slice[i])
	}
}
