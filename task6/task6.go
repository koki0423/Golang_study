package main
import "fmt"

func main() {
	slice:= [] string{"a","b","c","d","e"}

	for _, s := range slice {
		fmt.Println(s)
	}
}