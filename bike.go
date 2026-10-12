package main

import "fmt"

func main() {
	var r float64
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&r)

	area := (22/7) * r * r
	spin10 := area * 10
	fmt.Println("=================================================")
	fmt.Println("Keliling roda", area,)
	fmt.Println("Jarak Tempuh", spin10)

}
