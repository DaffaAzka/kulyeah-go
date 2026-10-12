package main

import "fmt"

func main() {
	var ferenheit int
	var celcius int
	fmt.Print("Masukan suhu dalam Ferenheit: ")
	fmt.Scanln(&ferenheit)

	celcius = (ferenheit - 32) * 5 / 9
	fmt.Println("=================================================")
	fmt.Println("Suhu dalam Ferenheit adalah", ferenheit, "derajat")
	fmt.Println("Suhu dalam Celcius adalah", celcius, "derajat")

}
