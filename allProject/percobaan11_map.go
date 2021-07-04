package main

import "fmt"

func main() {
	bank_soal := make(map[string]string) // map[key]value
	bank_soal["s1"] = "Siapa Presiden RI 1?"
	bank_soal["s1a"] = "Soekarno"
	bank_soal["s1b"] = "Soeharto"
	bank_soal["s1c"] = "BJ Habibie"
	bank_soal["s1d"] = "Gusdur"
	bank_soal["s1e"] = "Sajarwo"
	bank_soal["s1jawaban"] = "s1a"

	fmt.Println("Size Map :", len(bank_soal))
	fmt.Println(bank_soal)

	// menghapus elemen
	delete(bank_soal, "s1e")

	for k, v := range bank_soal {
		fmt.Printf("\nKey: %v, Val: %v", k, v)
	}

	// akses elemen dengan kata kunci / key = s1
	soal1, ok := bank_soal["s1"]
	fmt.Println("\n\nSoal 1 : ", soal1, "\nStatus : ", ok)

	// akses elemen dengan kata kunci / key = s2
	soal2, ok := bank_soal["s2"]
	fmt.Println("\nSoal 2 : ", soal2, "\nStatus : ", ok)
}
