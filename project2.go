package main

import "fmt"

func main() {
	var nama_depan, nama_belakang string
	fmt.Print("Masukkan Nama Depan Anda: ")
	fmt.Scan(&nama_depan)
	fmt.Print("Masukkan Nama Belakang Anda: ")
	fmt.Scan(&nama_belakang)
	fmt.Println("Hello, " + nama_depan + " " + nama_belakang + "!")
}
