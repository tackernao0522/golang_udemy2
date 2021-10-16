package main

import "fmt"

// マップ
// for

func main() {
	m := map[string]int{
		"Apple":  100,
		"Banana": 200,
	}

	for k, v := range m {
		fmt.Println(k, v)
		// Apple 100
		// Banana 200
	}

	for _, v := range m {
		fmt.Println(v)
		// 100
		// 200
	}

	for k := range m {
		fmt.Println(k)
		// Apple
		// Banana
	}
}
