package main

import "fmt"

func main() {
	var name string
	var nim string
	var class string
	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scanln(&name)
	fmt.Print("Masukkan NIM Anda: ")
	fmt.Scanln(&nim)
	fmt.Print("Masukkan Kelas Anda: ")
	fmt.Scanln(&class)

	fmt.Println("Perkenalkan saya adalah", name, ", salah satu mahasiswa Prodi S1-IF dari kelas", class, "dengan NIM", nim, ".")
}
