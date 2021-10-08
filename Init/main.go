package main

import "fmt"

// init

func init() {
	fmt.Println("init")
}

func init() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("Main")
}

// init関数を使うと先にinit関数が走るようになる
// init
// init2
// Main
