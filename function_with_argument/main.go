package main

import "fmt"

//関数
//関数を引数に取る関数

func CallFunction(f func()) {
	f() // 引数のf関数を実行
}

func main() {
	CallFunction(func() {
		fmt.Println("I'm a function")
	})
}
