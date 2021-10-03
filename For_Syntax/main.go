package main

import "fmt"

// for
// 繰り返し処理

func main() {
	// i := 0
	// for {
	// 	i++
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Println("Loop") // Loop Loop
	// }

	// point := 0
	// for point < 10 {
	// 	fmt.Println(point)
	// 	point++
	// }

	// for point := 0; point < 10; point++ {
	// 	fmt.Println(point)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	if i == 6 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// arr := [3]int{1, 2, 3}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i]) // 1 2 3
	// }

	// arr := [3]int{1, 2, 3}

	// for i, v := range arr {
	// 	fmt.Println(i, v) // 0 1 1 2 2 3
	// }

	// arr := [3]int{1, 2, 3}

	// for _, v := range arr {
	// 	fmt.Println(v) // 1 2 3
	// }

	// arr := [3]int{1, 2, 3}

	// for i := range arr { // インデックスのみ取る場合
	// 	fmt.Println(i) // 0 1 2
	// }

	// スライス
	// sl := []string{"Python", "PHP", "GO"}
	// for i, v := range sl {
	// 	fmt.Println(i, v) // 0 Python 1 PHP 2 GO
	// }

	// マップ
	// m := map[string]int{"apple": 100, "banana": 200} // keyはsting型 valueはint型
	// for k, v := range m {
	// 	fmt.Println(k, v) // apple 100 banana 200
	// }

	// m := map[string]int{"apple": 100, "banana": 200}
	// for _, v := range m {
	// 	fmt.Println(v) // 100 200
	// }

	m := map[string]int{"apple": 100, "banana": 200}
	for k := range m {
		fmt.Println(k) // apple banana
	}
}
