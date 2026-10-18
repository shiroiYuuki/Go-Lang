package main

import "fmt"

func main() {
	// deklare di awal
	var hewan string

	hewan = "Rubah"
	fmt.Println(hewan)

	// kalo ingin merubah harus dirubah ke tipe data yang sama
	hewan = "Kucing"
	fmt.Println(hewan)

	//  diklare saat di gunakan
	var foxWife = "fubuki"
	fmt.Println(foxWife)

	var IHSG = 5000
	fmt.Println(IHSG)

	// tanpa var
	negara := "indonesia"
	fmt.Println(negara)

	var (
		namaAwal  = "Shirakami"
		namaAkhir = "Fubuki"
	)

	fmt.Println(namaAwal)
	fmt.Println(namaAkhir)
}
