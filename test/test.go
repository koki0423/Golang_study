package main
import "fmt"

func main() {
    var a int
    var index int = 0
    for {
        fmt.Scan(&a)
        if a == 0 {
            return
        }
        fmt.Printf("index %d: %d\n", index, a)
        index++
    }
}
