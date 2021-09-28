package main

import "fmt"

// 定数

const Pi = 3.14 // 他のパッケージから呼び出したい時は頭文字を大文字

const (
	URL      = "http://xxx.co.jp"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const (
	c0 = iota // 連続する整数の連番を生成する
	c1
	c2
)

// var Big int = 923333678777876432 + 1

const Big = 923333678777876432 + 1

func main() {
	fmt.Println(Pi) // 3.14

	// Pi = 3 再代入できない
	// fmt.Println(Pi) エラー

	fmt.Println(URL, SiteName) // http://xxx.co.jp test

	fmt.Println(A, B, C, D, E, F) // 1 1 1 A A A

	// fmt.Println(Big - 1) // オーバーフローして定義できず

	fmt.Println(Big - 1)

	fmt.Println(c0, c1, c2) // 0 1 2
}
