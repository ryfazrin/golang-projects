package main

import "fmt"

func main() {
	// model 1
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}

	// model 2
	var j int = 0
	for j < 5 {
		fmt.Println("Nilai j = ", j)
		j++
	}

	// model 3
	for pos, char := range "SAJARWO" {
		fmt.Printf("character %#U start at byte position %d\n", char, pos)
	}

	// model 4
	for {
		j++
		if j <= 10 {
			fmt.Println("Nilai j = ", j)
		} else {
			break
		}
	}
}
