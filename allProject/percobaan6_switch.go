package main

import "fmt"

func main() {
	var huruf byte = 'A'
	switch huruf {
	case 'A':
		fmt.Println("Karakter A")
	case 'B':
		fmt.Println("Karakter B")
	case '?':
		fmt.Println("Karakter ?")
	}

	var nilai int32 = 50 // nilai silahkan diubah-ubah
	switch {
	case nilai > 90:
		fmt.Println("Nilai Anda A")
	case nilai > 70:
		fmt.Println("Nilai Anda B")
	case nilai > 60:
		fmt.Println("Nilai Anda C")
	default:
		fmt.Println("Nilai Anda D")
	}
}
