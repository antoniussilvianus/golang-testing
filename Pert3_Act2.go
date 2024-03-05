package main

import (
	"fmt"
)

func main() {
	var nilaiUTS, nilaiUAS float64
	fmt.Print("Masukkan nilai UTS: ")
	fmt.Scanln(&nilaiUTS)
	fmt.Print("Masukkan nilai UAS: ")
	fmt.Scanln(&nilaiUAS)

	totalNilai := (0.3 * nilaiUTS) + (0.7 * nilaiUAS)
	var grade string

	if totalNilai <= 20 {
		grade = "E"
	} else if totalNilai <= 40 {
		grade = "D"
	} else if totalNilai <= 60 {
		grade = "C"
	} else if totalNilai <= 80 {
		grade = "B"
	} else {
		grade = "A"
	}

	fmt.Printf("\nNilai UTS: %.2f\n", nilaiUTS)
	fmt.Printf("Nilai UAS: %.2f\n", nilaiUAS)
	fmt.Printf("Total Nilai: %.2f\n", totalNilai)
	fmt.Printf("Grade: %s\n", grade)
}
