package main

import (
	"fmt"
	"time"
)

// go goroutin

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 + time.Millisecond)
	}
}

func main() {
	go sub() // goと入れると並行処理が行われる
	go sub()

	for {
		fmt.Println("Main Loop")
		time.Sleep(150 * time.Millisecond)
	}
}
