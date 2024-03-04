package main

import "fmt"

func main() {
	var nilai1, nilai2 float64

	defer fmt.Println("-----Selesai------")

	fmt.Printf("Nilai 1 = ")
	fmt.Scan(&nilai1)

	fmt.Printf("Nilai 2 = ")
	fmt.Scan(&nilai2)

	fmt.Printf("Hasil = %.3f - %.3f = %.3f\n", nilai1, nilai2, nilai1-nilai2)

	if nilai2 != 0 {
		hasil := nilai1 / nilai2
		fmt.Printf("Hasil Dari Nilai 1 / Nilai 2 = %.3f\n", hasil)
	} else {
		fmt.Println("Tidak bisa melakukan pembagian karena nilai 2 adalah 0.")
	}
}
