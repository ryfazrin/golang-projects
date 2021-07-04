package main

import (
	"math"
	"fmt"
)

func main(){
	// contoh 1
	a:= math.Pow(2,8)
	fmt.Println("nilai a =", a)

	// contoh 2
	for i:=0;i<=8;i++ {
		b := math.Pow(2, float64(i))
		fmt.Println("nilai b =", b)
	}

	// contoh 3
	fmt.Println("nilai m.ceil  =", math.Ceil(3.5))
	fmt.Println("nilai m.floor =", math.Floor(3.5))
	fmt.Println("nilai m.round =", math.Round(3.5))
	fmt.Println("nilai m.round =", math.Round(3.4))

	// contoh 4
	n1 := 5.0
	n2 := 8.0
	fmt.Println("nilai min =", math.Min(n1,n2))
	fmt.Println("nilai max =", math.Max(n1,n2))
}
