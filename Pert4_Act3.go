package main

import "fmt"

func main() {

	var data = map[string]mahasiswa{
		"10122197": {
			nama:  "Antonius Silvianus",
			kelas: "2KA24",
		},
		"11111111": {
			nama:  "Young Lex Anjay",
			kelas: "1KA24",
		},
	}

	var search string
	fmt.Print("Masukkan NPM : ")
	fmt.Scanln(&search)

	var value, ok = data[search]

	if ok {
		fmt.Printf("Nama anda %s \n Kelas anda %s\n", value.nama, value.kelas)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

type mahasiswa struct {
	nama  string
	kelas string
}
