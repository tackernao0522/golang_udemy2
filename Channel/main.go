package main

import "fmt"

// channel (データの送受信を行うデータ型)
// 複数のゴルーチン間でのデータの受け渡しをする為に設計された構造。
// 宣言、操作

func main() {
	var ch1 chan int // 必ず定義

	// var ch2 <-chan int // 受信専用のチャネル

	// var ch3 chan<- int // 送信専用のチャネル

	ch1 = make(chan int) // チャネルの生成と初期化を行なって書き込みと読み込みができる状態

	ch2 := make(chan int)

	// 要領を調べる
	fmt.Println(cap(ch1)) // 0
	fmt.Println(cap(ch2)) // 0

	ch3 := make(chan int, 5) // 容量5のチャネルを生成 5はキューの数
	fmt.Println(cap(ch3))    // 5

	ch3 <- 1              // ch3(チャネル)に1のデータを送信したことになる
	fmt.Println(len(ch3)) // 1 一つデータを送ったので要素数が1

	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3)) // len 3

	// チャネルからデータを受信する キューは先入れ先出しになる
	i := <-ch3
	fmt.Println(i)               // 1 最初に入れたデータ
	fmt.Println("len", len(ch3)) // len 2

	i2 := <-ch3
	fmt.Println(i2)              // 2
	fmt.Println("len", len(ch3)) // len 1

	// 別の書き方
	fmt.Println(<-ch3)           // 3
	fmt.Println("len", len(ch3)) // len 0

	// バッファサイズを超えたデータを送った場合 (受信しないとデッドロックになるが一度受信することによってデッドロックを避けることができる)
	ch3 <- 1
	fmt.Println("len", len(ch3)) // len 1
	fmt.Println(<-ch3)           // 1
	fmt.Println("len", len(ch3)) // len 0
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6

}
