package main

import "fmt"

// マップ

func main() {
	var m = map[string]int{"A": 100, "B": 200} // keyがstring valueがintで指定したマップ
	fmt.Println(m)                             // map[A:100 B:200]

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2) // map[A:100 B:200]

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3) // map[1:A 2:B]

	m4 := make(map[int]string)
	fmt.Println(m4) // map[]

	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4) // map[1:JAPAN 2:USA]

	fmt.Println(m["A"]) // 100
	fmt.Println(m4[2])  // USA
	fmt.Println(m4[3])  // 登録されていないのは初期値が返ってくる stringなので空文字が返ってくる

	_, ok := m4[3] // 取り出しに成功しているか判定をかけている
	if !ok {
		fmt.Println("error")
	}
	// fmt.Println(s, ok) // JAPAN false

	m4[2] = "US"
	fmt.Println(m4) // map[1:JAPAN 2:US]

	m4[3] = "CHINA"
	fmt.Println(m4) // map[1:JAPAN 2:US 3:CHINA]

	delete(m4, 3)   // 第一引数は削除したいmapで第二引数はキーを入れる
	fmt.Println(m4) // map[1:JAPAN 2:US]

	fmt.Println(len(m4)) // 要素数の2が出力される
}
