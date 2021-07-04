package main

import (
	"net/http"
	"io/ioutil"
)

// direktori tempat bekerja
var Dir_Name ="E:/PSMK/Project/Webserver1/"

func main(){
	// membuat file server
	dir_name :=Dir_Name+"res/"
	fileServer := http.FileServer(http.Dir(dir_name))
	http.Handle("/res/", http.StripPrefix("/res/", fileServer))

	// menghandle halaman utama
	http.HandleFunc("/", HomePage)

	//menjalankan server dengan domain/ip dengan port tertentu
	http.ListenAndServe("127.0.0.1:8081",nil)
}

// fungsi menangani halaman utama
func HomePage(w http.ResponseWriter, r *http.Request){
	data:= ReadHtmlFile(Dir_Name+"index.html")
	w.Write(data)
}

// File Reader (output dalam bentuk []byte)
func ReadHtmlFile(name string) []byte {
	data, err := ioutil.ReadFile(name)
	if err !=nil{
		panic(0)
	}
	return data
}
