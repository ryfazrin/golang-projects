package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main(){
	// contoh 1
	rand.Seed(time.Now().Unix())
	r1 := rand.Int()
	r2 := rand.Intn(10)
	fmt.Println("nilai r1 =", r1)
	fmt.Println("nilai r2 =", r2)
}
