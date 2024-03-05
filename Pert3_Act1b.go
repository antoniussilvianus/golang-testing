package main

import "fmt"

func main() {
	n := 12 // jumlah angka yang akan ditampilkan
	for i := 1; i <= n; i += 2 {
		fmt.Printf("Angka %d\n", i)
	}
}
