package main

import "fmt"

func main() {
	/*
		var i int = 1
		fl64 := float64(i)
		fmt.Println(fl64)               // 1
		fmt.Printf("i = %T\n", i)       // int
		fmt.Printf("fl64 = %T\n", fl64) // float64

		i2 := int(fl64)
		fmt.Println(i2)                 // 1
		fmt.Printf("fl64 = %T\n", fl64) // float64
		fmt.Printf("i2 = %T\n", i2)     // int

		fl := 10.5
		i3 := int(fl)
		fmt.Printf("i3 = %T\n", i3) // int
		fmt.Println(i3)             // 10 小数点以下は切り捨てられる
	*/

	/*
		var s string = "100"
		fmt.Printf("s = %T\n", s) // string

		i, _ := strconv.Atoi(s) // Atoiは文字列型から数値型に変換する関数 errのところを _にするともう１つの関数は使用しないの意味
			if err != nil {
				fmt.Println(err)
			}
	*/
	/*
		fmt.Println(i)            // 100
		fmt.Printf("i = %T\n", i) // int

		var i2 int = 200
		s2 := strconv.Itoa(i2)
		fmt.Println(s2)             // 200
		fmt.Printf("s2 = %T\n", s2) // string
	*/

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b) // [72 101 108 108 111 32 87 111 114 108 100]
	h2 := string(b)
	fmt.Println(h2) // Hello World
}
