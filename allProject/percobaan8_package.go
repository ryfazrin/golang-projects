package main

import (
	"fmt"
	"pustakasaya"
)

func main() {
	jumlah := pustakasaya.Penjumlahan(3,5)
	fmt.Println("Hasil Penjumlahan = ", jumlah)

	perkalian := pustakasaya.Perkalian(3,5)
	fmt.Println("Hasil Perkalian = ", perkalian)

	pustakasaya.DeteksiBilanganNegatif(-19)
	
}
