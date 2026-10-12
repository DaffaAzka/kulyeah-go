package main

import "fmt"

func main() {
	var tinggi int
	fmt.Print("Masukkan tinggi air dalam meter: ")
	fmt.Scan(&tinggi)

	volume := 4 * 2 * tinggi
	berat := volume * 1000

	fmt.Println("Volume Air", volume)
	fmt.Println("Berat Air", berat)
}