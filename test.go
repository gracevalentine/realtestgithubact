package main

import (
	"fmt"
)

// nomor 1C
var (
	var_1 float32 = 20.5 // yang ini private, ditandai dgn huruf bukan kapital pd awal var
	Var_2 float3  = 20.5 // yang ini public, ditandai dgn huruf kapital pd awal var
	num   int
)

func main() {
	fmt.Prin("anjing jalan")
	for i := 1; i <= num; i++ {
		fmt.Print("*")
	}
}
