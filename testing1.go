package main

import "fmt"

var nilai1, nilai2 float64

func main() {
	defer fmt.Println("-----Selesai------")
	fmt.Print("Nilai 1 = ")
	fmt.Scan(&nilai1)
	fmt.Print("Nilai 2 = ")
	fmt.Scan(&nilai2)
	fmt.Printf("Hasil = %.3f - %.3f = %.3f\n", nilai1, nilai2, nilai1-nilai2)
	hasil := nilai1 / nilai2
	fmt.Printf("Hasil Dari Nilai 1 / Nilai 2 = %.3f\n", hasil)
}
