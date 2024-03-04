package main

import (
	"fmt"

)

const pi = 22 / 7

func main() {
	var jari int
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&jari)

	luas := int(pi * jari * jari)

	fmt.Print("Luas lingkaran = ", luas)
}
