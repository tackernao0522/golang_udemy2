package main

import "fmt"

func main() {
	var arr1 [3]int
	fmt.Println(arr1)        // [0 0 0]
	fmt.Printf("%T\n", arr1) // [3]int

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2) // [A B ]

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3) // [1 2 3]

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)        // [C D]
	fmt.Printf("%T\n", arr4) // [2]string

	fmt.Println(arr1[0])   // 0
	fmt.Println(arr2[0])   // A
	fmt.Println(arr2[1])   // B
	fmt.Println(arr2[2])   // (空)
	fmt.Println(arr2[2-1]) // B

	arr2[2] = "C"
	fmt.Println(arr2) // [A B C]

	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5) // エラー

	fmt.Println(len(arr1)) // 3
}
