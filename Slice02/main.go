package main

import "fmt"

// スライス
// append make len cap

func main() {
	sl := []int{100, 200}
	fmt.Println(sl) // [100, 200]

	sl = append(sl, 300)
	fmt.Println(sl) // [100 200 300]

	sl = append(sl, 400, 500, 600)
	fmt.Println(sl) // [100 200 300 400 500 600]

	sl2 := make([]int, 5)
	fmt.Println(sl2) // [0 0 0 0 0]

	fmt.Println(len(sl2)) // 5

	fmt.Println(cap(sl2)) // 5 capacity(容量)

	sl3 := make([]int, 5, 10) // 第二引数は要素数 第三引数はcapacity

	fmt.Println(len(sl3)) // 5

	fmt.Println(cap(sl3)) // 10

	sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)

	fmt.Println(len(sl3)) // 12

	fmt.Println(cap(sl3)) // 20
}
