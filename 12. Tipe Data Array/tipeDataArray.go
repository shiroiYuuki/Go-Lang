package main

import "fmt"

func main() {
	var name [3]string // untuk mendeklarasikan jumlah array yang dimana di sini adalah 3

	name[0] = "Tokino"
	name[1] = "Mary"
	name[2] = "Wife"

	fmt.Println("Nama Semuanya ", name)
	fmt.Println("Nama 0 adalah: ", name[0])
	fmt.Println("Nama 1 adalah: ", name[1])
	fmt.Println("Nama 2 adalah: ", name[2])

	var nilai = [3]int{ // cara lain membuat array
		90,
		67,
		45,
	}

	fmt.Println("Nilai semuanya ", nilai)
	fmt.Println("Nilai 0 adalah: ", nilai[0])
	fmt.Println("Nilai 1 adalah: ", nilai[1])
	fmt.Println("Nilai 2 adalah: ", nilai[2])

	fmt.Println(len(name))
	fmt.Println(len(nilai))

	var luas [13]int64     // data kosong
	fmt.Println(len(luas)) // yang di hitung adalah panjang data bukan jumlah isi data
}
