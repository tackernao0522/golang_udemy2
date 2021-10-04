package main

import "fmt"

// switch
// 型スイッチ

func anything(a interface{}) { // interface型の引数を渡す
	// fmt.Println(a) // aaa 1 様々な型を入れられる。但し計算はできない
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?") // aaa!?
	case int:
		fmt.Println(v + 1000) // 1001
	}
}
func main() {
	anything("aaa") // aaa!?
	anything(1)     // 1001

	var x interface{} = 3
	// 型アサーション
	i := x.(int)       // interface型をint型に復元している
	fmt.Println(i + 2) // 5 計算できる
	// fmt.Println(x + 2) // xはinterface型のままの為、計算できない。エラーになる

	// xはint型なのでfloatは復元できない エラー
	// f := x.(float64)
	// fmt.Println(f)

	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64) // 0 false runtimeエラーにならない 変換に失敗するとfalseと返ってくる

	i, isInt := x.(int)
	fmt.Println(i, isInt) // 3 trueと返ってくる

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt { // x = 3の isIntがtrueであれば
		fmt.Println(i, "x is Int") // 3 x is Int
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I don't know")
	}

	switch x.(type) {
	case int:
		fmt.Println("int") // int
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int") // 3 int
	case string:
		fmt.Println(v, "string")
	}
}
