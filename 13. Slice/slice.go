package main // Menandakan bahwa program ini bisa dijalankan langsung.

import "fmt" // Mengimpor fmt untuk memakai fungsi Println.

func main() { // Fungsi utama; program mulai berjalan dari sini.
	// Membuat array berisi 5 nama hewan.
	// Tanda [...] membuat Go menghitung jumlah elemennya secara otomatis.
	var hewan = [...]string{
		"Fox",   // Indeks 0
		"Wolf",  // Indeks 1
		"Cat",   // Indeks 2
		"dog",   // Indeks 3
		"Eagle", // Indeks 4
	}

	// Membuat slice dari indeks 2 sampai sebelum indeks 4.
	// Jadi, slice1 berisi hewan[2] dan hewan[3]: "Cat" dan "dog".
	var slice1 = hewan[2:4]

	fmt.Println(slice1)      // Mencetak: [Cat dog]
	fmt.Println(len(slice1)) // Panjangnya 2 karena ada 2 elemen di dalam slice.
	fmt.Println(cap(slice1)) // Kapasitasnya 3: mulai dari indeks 2, tersisa 3 tempat di array.

	// Jika diaktifkan, nilai hewan[4] berubah menjadi "Diubah".
	// slice1 tidak menampilkan perubahan ini karena indeks 4 tidak termasuk di dalam slice1.
	// hewan[4] = "Diubah"
	// fmt.Println(slice1)

	// Jika diaktifkan, slice1[0] berubah menjadi "Mei Lagi".
	// Karena slice1 berbagi penyimpanan dengan array hewan,
	// nilai hewan[2] juga akan berubah.
	// slice1[0] = "Mei Lagi"
	// fmt.Println(hewan)

	// Membuat slice dari indeks 2 sampai akhir array.
	// slice2 berisi: ["Cat", "dog", "Eagle"].
	var slice2 = hewan[2:]

	// Mencetak kapasitas slice2, yaitu 3.
	// Mulai dari indeks 2 sampai akhir array ada 3 elemen.
	fmt.Println(cap(slice2))

	// Mencetak seluruh isi array hewan.
	fmt.Println(hewan)

	// Membuat slice string dengan panjang 2 dan kapasitas 4.
	// Panjang 2 berarti indeks 0 dan 1 bisa langsung diakses.
	// Kapasitas 4 berarti slice dapat diperluas hingga 4 elemen
	// sebelum Go perlu menyediakan penyimpanan baru.
	newSlice := make([]string, 2, 4)

	// Mengisi elemen pertama, yaitu indeks 0.
	newSlice[0] = "Tokino"

	// Mengisi elemen kedua, yaitu indeks 1.
	newSlice[1] = "Mary"

	// Mencetak isi newSlice: [Tokino Mary].
	fmt.Println(newSlice)

	// Mencetak jumlah elemen yang sedang ada: 2.
	fmt.Println(len(newSlice))

	// Mencetak kapasitas slice: 4.
	fmt.Println(cap(newSlice))

	// Membuat slice baru dengan panjang dan kapasitas yang sama seperti newSlice.
	copySlice := make([]string, len(newSlice), cap(newSlice))

	// Menyalin elemen dari newSlice ke copySlice.
	// Keduanya memiliki isi yang sama, tetapi penyimpanannya terpisah.
	copy(copySlice, newSlice)

	// Mencetak isi copySlice: [Tokino Mary].
	fmt.Println(copySlice)

	// Membuat array int dengan panjang tetap 5.
	iniArray := [5]int{1, 2, 3, 4, 5}

	// Membuat slice int berisi angka yang sama.
	// Panjang slice dapat berubah, misalnya dengan menggunakan append.
	iniSlice := []int{1, 2, 3, 4, 5}

	// Mencetak isi array.
	fmt.Println(iniArray)

	// Mencetak isi slice.
	fmt.Println(iniSlice)
} // Menutup fungsi main.
