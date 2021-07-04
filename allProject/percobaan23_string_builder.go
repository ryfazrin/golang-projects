package main

import (
	"strings"
	"fmt"
)
func main(){
	// contoh 1
	sb := strings.Builder{}

	sb.WriteString("nama saya budi")

	fmt.Println(sb)
	fmt.Println(sb.String())
	fmt.Println(sb.Len())

	// mengosongkan strings builder
	sb.Reset()
	fmt.Println(sb.String())

	// contoh 2
	sr:= strings.NewReader("nama saya budi")

	// membuat wadah penyimpanan
	r := make([]byte, 4)

	//mengambil string dan menempatkan ke wadah, mulai dari offset
	sr.ReadAt(r,5)

	fmt.Println(string(r))
}
