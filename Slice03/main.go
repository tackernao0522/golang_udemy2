package main

import "fmt"

// スライス
// copy

func main() {
	// sl := []int{100, 200}
	// sl2 := sl

	// sl2[0] = 1000

	// fmt.Println(sl) // [1000 200]  sl2とslは同じメモリーアドレスを使っているため、slも書き換えられてしまう 参照型の場合は同じ場所をさすと言うこと。

	// var i int = 10
	// i2 := i
	// i2 = 100
	// fmt.Println(i, i2) // 10 100

	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)
	fmt.Println(sl2)    // [0 0 0 0 0]
	n := copy(sl2, sl)  // sl2にslをコピー
	fmt.Println(n, sl2) // 5 [1 2 3 4 5] nはコピーに成功した数になる
}
