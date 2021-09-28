package main

import "fmt"

// 算術演算子

func main() {
	fmt.Println(1 + 2)        // 3
	fmt.Println("ABC" + "DE") // ABCDE

	fmt.Println(5 - 1) // 4

	fmt.Println(5 * 4) // 20

	fmt.Println(6 / 3) // 2

	fmt.Println(9 % 4) // 1

	n := 0
	n += 4
	fmt.Println(n) // 4
	n++
	fmt.Println(n) // 5
	n--
	fmt.Println(n) // 4

	s := "ABC"
	s += "DEF"
	fmt.Println(s) // ABCDEF
}
