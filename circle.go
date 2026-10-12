package main

import "fmt"

func main() {
	var radius float64
	const pi = 3.14

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&radius)

	area := pi * radius * radius
	fmt.Printf("Luas lingkaran dengan jari-jari %.2f adalah %.2f\n", radius, area)
}
