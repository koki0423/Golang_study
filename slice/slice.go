package main

import "fmt"

func main() {
	// make(スライス型, 初期個数, 初期容量)
	slice := make([]int, 0, 2)
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))
	if len(slice) > 0 {
		fmt.Printf("Address: %p ~ %p\n", &slice[0],&slice[len(slice)-1])
	}

	slice = append(slice, 1)
	slice = append(slice, 2)
	fmt.Printf("Length: %d, Capacity: %d \n", len(slice), cap(slice))
	fmt.Printf("Address: %p ~ %p\n", &slice[0],&slice[len(slice)-1])

	//規定サイズ超えたら自動で拡張される
	slice = append(slice, 3)
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))
	fmt.Printf("Address: %p ~ %p\n", &slice[0],&slice[len(slice)-1])
}
