package main

import (
	"fmt"
	"os"
)

// defer

func TestDefer() { // 関数が終了時に deferで定義した関数が実行されるので先にSTARTが出力される
	defer fmt.Println("END")
	fmt.Println("START")
	// START
	// END
}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func main() {
	TestDefer()

	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()

	// START
	// END
	// 1
	// 2
	// 3

	RunDefer()
	// START
	// END
	// 3 後からdeferで登録された値が最初に出力される
	// 2
	// 1

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("Hello"))
}
