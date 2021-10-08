package main

import "fmt"

// スライス
// 宣言、操作

func main() {
	var sl []int
	fmt.Println(sl) // []

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2) // [100 200]

	sl3 := []string{"A", "B"}
	fmt.Println(sl3) // [A B]

	sl4 := make([]int, 5)
	fmt.Println(sl4) // [0 0 0 0 0]

	sl2[0] = 1000
	fmt.Println(sl2) // [1000 200]

	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5) // [1 2 3 4 5]

	fmt.Println(sl5[0]) // 1

	fmt.Println(sl5[2:4]) // [3 4] index0とindex4の手前

	fmt.Println(sl5[:2]) // [1 2] index2の手前まで

	fmt.Println(sl5[2:]) // [3 4 5] index2からあとは全部

	fmt.Println(sl5[:]) // [1 2 3 4 5] 全部

	fmt.Println(sl5[len(sl5)-1]) // 5 一番最後の要素を取る

	fmt.Println(sl5[1 : len(sl5)-1]) // [2 3 4] 一番目と最後を除く
}
