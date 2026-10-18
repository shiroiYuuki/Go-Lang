package main

import "fmt"

func main() {
	// kelebihan constan jika tidak di pakai, tidak akan error
	// constan wajib ada "consta"
	const nama string = "Malfoid"
	const nama2 = "Shirakami"
	const value = 5000

	// bisa juga di diklare sekaligus begini
	const (
		umur  int = 5000
		nama3     = "Whisper"
	)

	fmt.Println(nama)
	fmt.Println(nama2)
	fmt.Println(value)

	fmt.Println(umur)
	fmt.Println(nama3)
}
