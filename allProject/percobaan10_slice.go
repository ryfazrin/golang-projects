package main

import "fmt"

func main(){
	// contoh 1
	data1 := make([]int, 5)
	for i := 0; i < 5; i++ {
		data1[i] = i
		fmt.Println(data1[i])
	}

	// contoh 2
	data2 := []int{ 1, 2, 3, 4, 5, 6 }

	split1 := data2[:] // [1 2 3 4 5 6]
	split2 := data2[2:5] // [3 4 5]
	split3 := data2[3:] // [4 5 6]
	split4 := data2[:3] // [1 2 3]

	fmt.Println("data2   = ",data2)
	fmt.Println("split1  = ",split1)
	fmt.Println("split2  = ",split2)
	fmt.Println("split3  = ",split3)
	fmt.Println("split4  = ",split4)

	// contoh 3
	data2 = append(data2, 7)
	fmt.Println("\ndata2 append   = ",data2)

	// contoh 4
	data3 := make([]int, 5)
	copy(data3, data2)         	// ganti dengan: data3 = data2
	fmt.Println("\ncopy data2 to data3   = ",data3)

	data2[0]=99
	fmt.Println("\ncopy data2 to data3   = ",data3)
}
