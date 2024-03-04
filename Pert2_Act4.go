package main

import "fmt"

func main() {
	var nilai1, nilai2, nilai3, nilai4, nilai5 float64
	fmt.Println("Masukkan nilai 1, nilai 2, nilai 3, nilai 4, dan nilai 5:")
	fmt.Scan(&nilai1, &nilai2, &nilai3, &nilai4, &nilai5)

	hasil := (nilai1 * nilai2) + nilai3 - (nilai4 / nilai5)
	fmt.Printf("Hasil = (%.3f * %.3f) + %.3f - (%.3f / %.3f) = %.3f\n", nilai1, nilai2, nilai3, nilai4, nilai5, hasil)
}
