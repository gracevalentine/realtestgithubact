package main

import (
	"fmt"
)

var (
	var1 float32 = 20.5 // yang ini private, ditandai dgn huruf bukan kapital pd awal var
	Var2 float32 = 20.5 // yang ini public, ditandai dgn huruf kapital pd awal var
	num  int
)

func main() {
	fmt.Print("akhirnya anjing jalan")
	for i := 1; i <= num; i++ {
		fmt.Print("*")
	}
}
