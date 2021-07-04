package pustakasaya

import "fmt"

func DeteksiBilanganNegatif(a int) {
	if a < 0 {
		fmt.Println(a, "adalah bilangan negatif")
	} else {
		fmt.Println(a, "adalah bilangan positif")
	}

}
