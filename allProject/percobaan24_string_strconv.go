package main

import (
	"strconv"
	"fmt"
)

func main(){
	// konversi dari string ke tipe lain
	a, err := strconv.ParseBool("true")
	b, err := strconv.ParseInt("42", 10, 64)
	c, err := strconv.ParseFloat("3.1415", 64)

	d, err:= strconv.Atoi("42")
	if err!=nil{
		panic(0)
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// konversi dari int to string
	s1 := strconv.Itoa(42)
	fmt.Println(s1)


	s2:= "nilai = "+ strconv.Itoa(7)
	fmt.Println(s2)

	s3:= strconv.AppendInt([]byte("nilai = "),7,10)
	fmt.Println(string(s3))
}
