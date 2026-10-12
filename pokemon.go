package main

import "fmt"

func main() {
	var ticket int
	isBuy := ticket >= 1500
	fmt.Print("Masukan jumlah tiket AEON: ")
	fmt.Scan(&ticket)
	fmt.Print(isBuy)
}
