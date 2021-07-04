package main

import "fmt"

func main()  {
  // variabel statis deklarasi
  var x int
  x = 10
  fmt.Println(x);
  // mencetak menggunakan Printf
  fmt.Printf("x memiliki tipe data %t\n", x)

  // variabel dinamis deklarasi
  z := "ryan"
  fmt.Println(z)

  i := 5
  f := 10
  fmt.Println(i+f)

  // penggunaan function penjumlahan
  jumlah := penjumlahan(1, 2)
  fmt.Println("hasil dari fungsi adalah", jumlah)
}

func penjumlahan(a int, b int) (int) {
  hasil := a+b
  return hasil
}
