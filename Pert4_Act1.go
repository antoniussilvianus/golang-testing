package main

import "fmt"

func main() {
	var gajisekarang, ekspektasigaji int

	fmt.Print("Masukkan gaji sekarang : ")
	fmt.Scan(&gajisekarang)

	fmt.Print("Masukkan gaji yang diinginkan : ")
	fmt.Scan(&ekspektasigaji)

	ekspektasigaji = naikangaji(gajisekarang, ekspektasigaji)
	fmt.Printf("\nEkspektasi gaji anda : %d\n", ekspektasigaji)
}

func naikangaji(gajisekarang, ekspektasigaji int) int {
	return ekspektasigaji
}
