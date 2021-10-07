package main

import "fmt"

// panic recover プログラムを強制的に終了させてしまう機能

func main() {
	defer func() { // recoverはruntimeエラーから復帰させる機能
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	panic("runtime error")
	fmt.Println("Start")
	// panic: runtime error
}

// 使用される場面はない
