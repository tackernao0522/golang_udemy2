package main

import "fmt"

// スライス
// 可変長引数

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}
func main() {
	fmt.Println(Sum(1, 2, 3))                       // 6
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) // 55
	fmt.Println(Sum())                              // 0

	sl := []int{1, 2, 3}
	fmt.Println(Sum(sl...)) // 6 スライスを実引数にして渡せる
}
