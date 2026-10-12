package main

import "fmt"

func convertion() {
	var celsius float64
	fmt.Print("Masukkan suhu dalam Celsius: ")
	fmt.Scanln(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Printf("%.2f derajat Celsius sama dengan %.2f derajat Fahrenheit\n", celsius, fahrenheit)
}
