package main

import "fmt"

func main() {
	var kursus = []string{"dbms", "server os", "networking", "web", "desktop", "erp"}
	fmt.Println("Slice kursus:")
	printSliceInfo(kursus)

	var kursusSaya = kursus[1:5]
	kursusSaya = append(kursusSaya, "Manajemen Informatika")
	fmt.Println("\nSlice kursus saya:")
	printSliceInfo(kursusSaya)
}

func printSliceInfo(s []string) {
	fmt.Printf("Panjang: %d, Kapasitas: %d, Isi: %v\n", len(s), cap(s), s)
}
