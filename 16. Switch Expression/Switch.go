package main

import "fmt"

func main() {

	name := "Mary"

	switch name {
	case "Tokino":
		fmt.Println("Hello Tokino")
		fmt.Println("Hello Tokino")
	case "Mary":
		fmt.Println("Hello Mary")
		fmt.Println("Hello Mary")
	default:
		fmt.Println("Kenalan Donk")
		fmt.Println("Kenalan Donk")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar panjang")
	}
}
