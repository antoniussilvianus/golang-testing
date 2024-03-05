package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Angka %d adalah genap\n", i)
		} else {
			fmt.Printf("Angka %d adalah ganjil\n", i)
		}
	}
}
