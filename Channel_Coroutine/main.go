package main

import (
	"fmt"
	"time"
)

// Channel
// チャネルとゴルーチン

func receiver(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// fmt.Println(<-ch1) // デッドロック

	go receiver(ch1)
	go receiver(ch2)

	i := 0
	for i < 100 {
		ch1 <- i
		ch2 <- i
		time.Sleep(50 * time.Millisecond)
		i++
	}
}
