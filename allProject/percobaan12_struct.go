package main

import "fmt"

type Siswa struct {
	nama   string
	nis    int
	berat  float32
	tinggi float32
}

func main() {
	s := new(Siswa)

	s.nama = "Budi"
	s.nis = 701
	s.berat = 50.2
	s.tinggi = 165

	fmt.Println("Nama Siswa  \t: ", s.nama)
	fmt.Println("Nomor Induk \t: ", s.nis)
	fmt.Println("Berat Badan \t: ", s.berat)
	fmt.Println("Tinggi Badan \t: ", s.tinggi)
}
