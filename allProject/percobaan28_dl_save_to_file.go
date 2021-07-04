package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
)

func main(){
	DownloadPageAndSaveToFile("http://merdeka.com","E:/PSMK/halaman_utama.html")

	// data yang dibaca dalam bentuk []byte
	data:=ReadFromFile("E:/PSMK/halaman_utama.html")

	// data dikonversi dari []byte ke dalam bentuk string
	fmt.Println(string(data))
}

// URL Downloader
func DownloadPageAndSaveToFile(url, name string){
	resp, err := http.Get(url)
	if err != nil {
		panic(0)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile(name, data, 0777)
}

// File Reader (output dalam bentuk []byte)
func ReadFromFile(name string) []byte {
	data, err := ioutil.ReadFile(name)
	if err !=nil{
		panic(0)
	}
	return data
}