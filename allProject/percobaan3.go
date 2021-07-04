package main

import "fmt"

func main() {
	// eksplisit
	var a = 1               // tanpa menuliskan tipe variable
	var b int = 2           // menyertakan tipe variable
	var c, d, e = 5, 10, 15 // lebih dari satu varible dalam satu baris
	var f float64 = 3.14    // inisialisasi variable float64

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c, d, e)
	fmt.Println(f)

	//implicit
	g := 100
	h := 1.1
	i, j := 4, 5.6
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i, j)
}
