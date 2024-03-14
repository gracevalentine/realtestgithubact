package main

import (
	"fmt"
)

// nomor 1C
var (
	var1 float32 = 20.5 // yang ini private, ditandai dgn huruf bukan kapital pd awal var
	Var2 float32 = 20.5 // yang ini public, ditandai dgn huruf kapital pd awal var
	num  int
)

func main() {
	fmt.Print("anjing jalan")
	for i := 1; i < num; i++ {
		fmt.Print("*")
	}
}
