package main

import (
	"fmt"
	"strings"
)

func main() {

	// char
	v1:='a'
	fmt.Println("char   = ",v1)
	fmt.Println("string = ",string(v1))

	v2:="aku"
	fmt.Println("v2[0] = ",v2[0])
	fmt.Println("v2[1] = ",v2[1])

	// konversi dari string ke byte
	v3:= []byte(v2)
	fmt.Println(v3)
	fmt.Println(string(v3))

	// konversi dari byte ke string
	v4:=[]byte{97, 98, 99, 100}
	fmt.Println(v4)
	fmt.Println(string(v4))


	a := "saya"
	b := "budi"
	c := "ani"
	fmt.Println(a, b, c)

	// contoh 1
	if a == b{
		fmt.Println(a, b)
	}
	if b == c {
		fmt.Println(b, c)
	}
	if b == "budi"{
		fmt.Println("var b = budi")
	}

	//contoh 2
	if strings.Compare(a,b) == 1{
		fmt.Println(strings.Compare(a,b))
	}
	if strings.Compare(b,a) == -1{
		fmt.Println(strings.Compare(b,a))
	}
	if strings.Compare(b,"budi") == 0{
		fmt.Println(strings.Compare(b,"budi"))
	}

	// memeriksa apakah terdapat kata tertentu di dalam string
	s := "nama saya budi"
	s1 := strings.Contains(s, "budi")
	s2 := strings.Contains(s, "bu")
	s3 := strings.Contains(s, "xi")
	fmt.Println(s1, s2, s3)

	// menghitung jumlah kata di dalam string
	fmt.Println(strings.Count(s, "a"))

	// membandingkan string tanpa melihat huruf besar atau kecil
	fmt.Println(strings.EqualFold("Go", "go"))
	fmt.Println(strings.EqualFold("Go", "Go"))

	// memilah kata menjadi slice by space
	s4 := strings.Fields(s)
	for i:=0;i<len(s4);i++{
		fmt.Println(i, s4[i])
	}
	for k,v:=range s4{
		fmt.Println(k, v)
	}

	// memilah kata menjadi array of slice berdasarkan separatornya
	s5:= strings.Split("aku pergi ke sekolah", " ")
	s6:= strings.Split("apel,jeruk,kedondong", ",")
	fmt.Println(s5)
	fmt.Println(s6)

	// manggabungkan array kata dengan separator
	fmt.Println(strings.Join(s4, ", "))


	// membandingkan pada awal kata
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasPrefix("Gopher", "pher"))

	// membandingkan pada akhir kata
	fmt.Println(strings.HasSuffix("Gopher", "Go"))
	fmt.Println(strings.HasSuffix("Gopher", "pher"))

	// mencari index posisi kata yang dicari
	fmt.Println(strings.Index("pelatihan", "latih"))
	fmt.Println(strings.Index("pelatihan", "psmk"))

	// mengubah kata yang terdapat di dalam string
	fmt.Println(strings.Replace("Pelatihan go lang psmk", "psmk","PSMK",1))

	// mengubah awal kata menjadi huruf kapital
	fmt.Println(strings.Title("Pelatihan go lang psmk"))

	// mengubah kata menjadi huruf kapital
	fmt.Println(strings.ToTitle("Pelatihan go lang psmk"))
	fmt.Println(strings.ToUpper("Pelatihan go lang psmk"))

	// mengubah kata menjadi huruf kecil
	fmt.Println(strings.ToLower("Pelatihan go lang psmk"))

	// menghapus awal dan/atau akhir kata berdasarkan katakunci tertentu
	fmt.Println(strings.Trim("SMK BISA  ", " "))
	fmt.Println(strings.Trim(",,SMK BISA,,,", ","))

	fmt.Println(strings.TrimLeft(",,SMK BISA,,,", ","))
	fmt.Println(strings.TrimRight(",,SMK BISA,,,", ","))

	// menghapus spasi pada awal dan akhir kata
	fmt.Println(strings.TrimSpace("  SMK  "))
}
