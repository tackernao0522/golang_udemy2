package main

import "fmt"

// 関数
// 関数を返す関数

func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function") // I'm a function
	}
}

func main() {
	f := ReturnFunc() // fには returnのfunc()が返っている
	f()
}
