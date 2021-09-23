package main

import "fmt"

func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl) // floatは環境依存ではないので普通に計算できる
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2 // float32はあまり使わない
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))
	// fl2 := 2.2
	// fmt.Println(float32(fl32) + float32(fl2))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf) // +Inf

	ninf := -1.0 / zero
	fmt.Println(ninf) // -Inf

	nan := zero / zero
	fmt.Println(nan) // NaN
}
