package main

import "fmt"

// 関数

// func Plus(x int, y int) int { // 返り値の型もいれる
// 	return x + y
// }

// func main() {
// 	i := Plus(1, 2)
// 	fmt.Println(i) // 3
// }

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	fmt.Println("No Return") // No Return
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i) // 3

	i2, i3 := Div(9, 4) // i2, _ := Div(9, 4)とするとi2のみ指定できる
	fmt.Println(i2, i3) // 2 1

	i4 := Double(1000)
	fmt.Println(i4) // 2000

	Noreturn()
}
