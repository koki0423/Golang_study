package main

import "fmt"

func main() {
    fmt.Print("整数 a を入力してください: ")
    var a int
    fmt.Scan(&a)
    fmt.Print("整数 b を入力してください: ")
    var b int
    fmt.Scan(&b)

    fmt.Println("a+b =", a+b)
    fmt.Println("a-b =", a-b)
    fmt.Println("a*b =", a*b)

    if b != 0 {
        fmt.Println("a/b =", a/b)
    } else {
        fmt.Println("エラー: 0で割ることはできません")
    }
}
