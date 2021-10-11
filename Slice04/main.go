package main

import "fmt"

// スライス
// for

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl) // [A B C]

	// for i, v := range sl {
	// 	fmt.Println(i, v)
	// 	// 0 A
	// 	// 1 B
	// 	// 2 C
	// }

	// for _, v := range sl {
	// 	fmt.Println(v)
	// 	// A
	// 	// B
	// 	// C
	// }

	// for i := range sl {
	// 	fmt.Println(i)
	// }
	// // 0
	// // 1
	// // 2

	// for i := range sl {
	// 	fmt.Println(i, sl[i])
	// }
	// // 0 A
	// // 1 B
	// // 2 C

	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
		// A
		// B
		// C
	}
}
