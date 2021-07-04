package main

import "fmt"

func main() {
	var i int = 21
	var j bool = true

	fmt.Printf("%v \n", i)   // menampilakan nilai i : 21
	fmt.Printf("%T \n", i)   // menampilkan tipe data dari variabel i
	fmt.Printf("%% \n")      // menampilkan tanda %
	fmt.Printf("%t \n\n", j) // menampilkan nilai boolean j : true

	fmt.Printf("%b \n", i)        // menampilkan nilai boolean j : true
	fmt.Printf("%c \n", '\u042F') // menampilkan unicode russia : Я (ya)
	fmt.Printf("%d \n", i)        // menampilkan nilai base 10 : 21
	fmt.Printf("%o \n", i)        // menampilkan nilai base 8  :25
	fmt.Printf("%x \n", 15)       // menampilkan nilai base 16 : f
	fmt.Printf("%X \n", 15)       // menampilkan nilai base 16 : F
	fmt.Printf("%U \n\n", 'Я')    // menampilkan unicode karakter Я : U+042F

	var k float64 = 123.456
	fmt.Printf("%f \n", k) // menampilkan nilai float : 123.456000
	fmt.Printf("%E \n", k) // menampilkan nilai float scientific : 1.234560E+02

}
