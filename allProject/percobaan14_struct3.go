package main

import (
	"fmt"
	"pustakasaya"
)

func main(){
	s := new(pustakasaya.Siswa)
	s.Nama = "Budi"
	s.Nis = 701
	s.Berat = 50.2
	s.SetTinggi(165)

	fmt.Println("Nama Siswa  \t: ", s.Nama)
	fmt.Println("Nomor Induk \t: ", s.Nis)
	fmt.Println("Berat Badan \t: ", s.Berat)
	fmt.Println("Tinggi Badan \t: ", s.Tinggi)
}
