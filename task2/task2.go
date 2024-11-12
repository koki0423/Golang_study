package main

import "fmt"

func main() {
	fmt.Print("文字列を入力: ")
	var str string
	fmt.Scan(&str)

	// 文字列をruneのスライスに変換
	runes := []rune(str)

	// runes スライスを逆順にする
	//配列中央で止まるようにすることで計算量を減らせる
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// runesを文字列に戻して出力
	fmt.Println("逆転した文字列:", string(runes))
}
