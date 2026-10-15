package main

import "fmt"

func main() {
	fmt.Println("Shirakami Fubuki")
	fmt.Println("Shirakami")
	fmt.Println("Fubuki") // ini adalah string
	fmt.Println("")       // kosong juga bisa

	fmt.Println(len("Shirakami Fubuki")) // menghitung panjang kalimat
	fmt.Println("Shirakami"[3])          // mencari urutan ke 3 di mulai dari 0
	fmt.Println("Fubuki"[0])             // yang di ambil itu kode ASCII/UTF-8 bukan index ke berapanya
}
