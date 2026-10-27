package main // Menandakan program ini bisa dijalankan langsung.

import "fmt" // Mengimpor fmt untuk menggunakan fungsi Println.

func main() { // Fungsi utama; program mulai berjalan dari sini.
	name := "Mary" // Membuat variabel name dan mengisinya dengan "Mary".

	// Memeriksa apakah name sama dengan "Tokino".
	if name == "Tokino" {
		fmt.Println("Hello Tokino") // Dijalankan jika name berisi "Tokino".
	} else if name == "Mary" {
		// Diperiksa jika kondisi pertama salah.
		fmt.Println("Hello Mary") // Dijalankan karena name berisi "Mary".
	} else {
		// Dijalankan jika kedua kondisi sebelumnya salah.
		fmt.Println("Hi, kenalan donk")
	}

	// length dibuat di dalam pernyataan if dan hanya bisa digunakan
	// di bagian if-else ini.
	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang") // Dijalankan jika panjang name lebih dari 5.
	} else {
		fmt.Println("Nama sudah benar") // Dijalankan jika panjang name 5 atau kurang.
	}
}
