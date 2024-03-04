package	main

import "fmt"
const pi = 22/7
var jari float64
var luas float64
func main() {
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&jari)
	luas = pi * jari * jari
	fmt.Print("Luas lingkaran = ",luas)
}