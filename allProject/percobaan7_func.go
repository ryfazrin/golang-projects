package main

import "fmt"

func main() {
	jumlah := Penjumlahan(3, 5)
	fmt.Println("Hasil Penjumlahan = ", jumlah)

	hasil_perkalian := perkalian(3, 5)
	fmt.Println("Hasil Perkalian = ", hasil_perkalian)
}

// penulisan nama fungsi huruf kapital
func Penjumlahan(a int, b int) int {
	c := a + b
	return c
}

// penulisan nama funsi huruf kecil
func perkalian(a int, b int) int {
	c := a * b
	return c
}
