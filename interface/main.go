package main

import "fmt"

func main() {
	var x interface{}
	fmt.Println(x) // nil

	x = 1
	fmt.Println(x) // 1

	x = 3.14
	fmt.Println(x) // 3.14

	x = "A"
	fmt.Println(x) // "A"

	x = [3]int{1, 2, 3}
	fmt.Println(x) // [1, 2, 3]

	// x = 2
	// fmt.Println(x + 3) // interface型とint型の計算はできない
}
