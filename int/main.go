package main

import "fmt"

func main() {
	var i int = 100

	var i2 int64 = 200

	fmt.Println(i + 50)

	// fmt.Println(i + i2)

	fmt.Printf("%T\n", i2) // i2の型を表示

	fmt.Printf("%T\n", int32(i2)) // i2をint32に型を変換

	fmt.Println(i + int(i2)) // 計算できるようになる

	var i3 int8 = 50

	var i4 int16 = 499

	fmt.Println(int(i3) + int(i4))
}
