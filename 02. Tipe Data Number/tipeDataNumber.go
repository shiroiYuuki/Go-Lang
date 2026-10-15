package main

import "fmt"

func main() {
	fmt.Println("Minus Satu: ", -1)      // Integer
	fmt.Println("Dua: ", 2)              // Unsigned Integer
	fmt.Println("Tiga Koma Lima: ", 3.5) // Float
}

/*
===================================================
 RINGKASAN TIPE DATA NUMBER & RANGE DI GO-LANG
===================================================

1. Integer (Bilangan Bulat, Mendukung Negatif)
   - int8  : -128  s/d  127
   - int16 : -32768  s/d  32767
   - int32 : -2147483648  s/d  2147483647
   - int64 : -9223372036854775808  s/d  9223372036854775807

2. Unsigned Integer (Bilangan Bulat, Hanya Positif)
   - uint8  : 0  s/d  255
   - uint16 : 0  s/d  65535
   - uint32 : 0  s/d  4294967295
   - uint64 : 0  s/d  18446744073709551615

3. Floating Point (Bilangan Desimal)
   - float32 : Desimal presisi 32-bit (Gunakan titik, cth: 3.5)
   - float64 : Desimal presisi 64-bit (Lebih presisi)

4. Complex Number (Untuk hitungan matematis/statistik)
   - complex64, complex128

5. Alias (Nama Lain yang Sering Dipakai)
   - byte : Alias untuk uint8 (0 s/d 255)
   - rune : Alias untuk int32 (-2 Milyar s/d 2 Milyar)
   - int  : Menyesuaikan OS (int32 di 32-bit, int64 di 64-bit)
   - uint : Menyesuaikan OS (uint32 di 32-bit, uint64 di 64-bit)

TIPS:
- Pilih range tipe data yang paling pas dengan data yang disimpan agar hemat memori (contoh: nilai ujian cukup pakai uint8).
===================================================
*/
