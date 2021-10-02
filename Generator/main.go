package main

import "fmt"

// 関数
// ジェネレーターの実装

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	ints := integers()
	fmt.Println(ints()) // 1
	fmt.Println(ints()) // 2
	fmt.Println(ints()) // 3
	fmt.Println(ints()) // 4

	otherints := integers()
	fmt.Println(otherints()) // 1
	fmt.Println(otherints()) // 2
	fmt.Println(otherints()) // 3
}
