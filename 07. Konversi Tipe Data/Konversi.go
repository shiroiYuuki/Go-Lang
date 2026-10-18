package main

import "fmt"

func main() {
	var nilai32 int32 = 32768 // limit di int 16 adalah 32,767
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai64) // ini akan memicu Integer Overflow jika angkanya melebihi batas

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
}
