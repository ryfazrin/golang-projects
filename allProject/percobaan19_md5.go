package main

import (
	"crypto/md5"
	"fmt"
)

func main(){
	output:= md5.Sum([]byte("pass123456"))
	fmt.Printf("%x", output)
}
