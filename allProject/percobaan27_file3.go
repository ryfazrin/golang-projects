package main

import (
	"os"
	"fmt"
	"bufio"
)

func main(){
	data:=[]string{}
	data=append(data,"aku")
	data=append(data,"pergi")
	data=append(data,"ke")
	data=append(data,"sekolah")

	CreateFile("E:/PSMK/data3.txt")
	AppendFileWithData("E:/PSMK/data3.txt", data)

	data_dari_file:=ReadFileWithData("E:/PSMK/data3.txt")
	fmt.Println(data_dari_file)
	for k,v:=range data_dari_file{
		fmt.Println(k,v)
	}
}

func CreateFile(name string){
	f, err := os.Create(name)
	if err !=nil{
		panic(0)
	}
	defer f.Close()
}

func AppendFileWithData(name string, d []string, ){
	f, err := os.OpenFile(name, os.O_APPEND,0777)
	if err !=nil{
		panic(0)
	}
	defer f.Close()

	for k,v:=range d{
		n,e:=f.WriteString(v+"\n")
		fmt.Println(k,v,n,e)
	}
}

func ReadFileWithData(name string)[]string{
	f, err := os.Open(name)
	if err !=nil{
		panic(0)
	}
	defer f.Close()

	data:=[]string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data=append(data,scanner.Text())
	}
	return data
}
