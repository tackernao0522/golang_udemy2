package main

import "fmt"

func main() {
	byteA := []byte{72, 73}
	fmt.Println(byteA)         // [72 73]
	fmt.Println(string(byteA)) // HI
	c := []byte("HI")
	fmt.Println(c)         // [72 73]
	fmt.Println(string(c)) // HI
}
