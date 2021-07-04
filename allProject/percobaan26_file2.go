package main

import (
	"os"
	"bufio"
	"fmt"
	"io"
)

func main(){
	CreateFile2()
	ReadFile2()
	ReadFile3()
	AppendFile()
}

func CreateFile2(){
	f, err := os.Create("E:/PSMK/data2.txt")
	if err !=nil{
		panic(0)
	}
	defer f.Close()

	f.WriteString("SMK INDONESIA\nBISA\nHEBAT")
}

func ReadFile2(){
	f, err := os.Open("E:/PSMK/data2.txt")
	if err !=nil{
		panic(0)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}
}

func ReadFile3(){
	f, err := os.Open("E:/PSMK/data2.txt")
	if err !=nil{
		panic(0)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func AppendFile(){
	f, err := os.OpenFile("E:/PSMK/data2.txt", os.O_APPEND,0777)
	if err !=nil{
		panic(0)
	}
	defer f.Close()

	n,e:=f.WriteString("\nBERHASIL")
	fmt.Println(n,e)
}