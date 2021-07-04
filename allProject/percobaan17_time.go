package main

import (
	"fmt"
	"time"
)

func main() {
	// contoh 1
	t := time.Now()
	fmt.Println(t)

	tahun, bulan, tanggal := t.Date()
	jam, menit, detik := t.Clock()
	fmt.Println("Date = ", tahun, bulan, tanggal )
	fmt.Println("Clock = ", jam, menit, detik)

	// contoh 2
	elapsed:= time.Now()
	for i:=0;i<10;i++{
		fmt.Println(time.Now())
		time.Sleep(1*time.Millisecond)
	}
	fmt.Println(time.Since(elapsed))
}
