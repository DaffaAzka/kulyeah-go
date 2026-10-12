package main

import "fmt"

func main() {
    var celsius float64
    fmt.Scan(&celsius)

    reamur := celsius * 4 / 5
    fahrenheit := celsius * 9 / 5 + 32
    kelvin := celsius + 273.15

    fmt.Printf("Derajat Reamur: %.0f\n", reamur)
    fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
    fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)
}