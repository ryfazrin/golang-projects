package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	CreateFile1()
	ReadFile1()
}

func ReadFile1()  {
	data, err := ioutil.ReadFile("E:/PSMK/data.txt")
	if err !=nil{
		panic(0)
	}
	fmt.Println(string(data))
}

func CreateFile1()  {
	data := []byte("SMK\nBISA\nHEBAT")
	err := ioutil.WriteFile("E:/PSMK/data.txt", data, 0777)
	if err !=nil{
		panic(0)
	}
}
