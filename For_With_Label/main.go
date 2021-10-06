package main

import "fmt"

// ラベル付きfor

func main() {
	// Loop:
	// 	for {
	// 		for {
	// 			for {
	// 				fmt.Println("STAR")
	// 				break Loop
	// 			}
	// 			// ①
	// 			fmt.Println("処理をしない")
	// 		}
	// 		// ②
	// 		fmt.Println("処理をしない")
	// 	}
	// 	fmt.Println("END") // ①と②の処理はされずにここが実行される STAR END

Loop:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue Loop // Loop: から このLoopを回せば最下部のfmtは実行されなくなる
			}
			fmt.Println(i, j, i*j)
			// 0 1 0
			// 1 1 1
			// 2 1 2
		}
		fmt.Println("処理をしない")
	}
}
