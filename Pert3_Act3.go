package main

import "fmt"

func main() {
	var i int

	for j := 1; j <= 10; j++ {
		fmt.Printf("Input nilai %d : ", j)
		fmt.Scanln(&i)

		if i <= 10 {
			if i%2 == 0 {
				fmt.Printf("%d adalah Angka Genap\n", i)
			} else {
				fmt.Printf("%d adalah Angka Ganjil\n", i)
			}
		} else {
			fmt.Println("Pertanyaan selesai, karena angka I yang diinput sudah melebihi dari 10. Terimakasih")
			break
		}
	}
}
