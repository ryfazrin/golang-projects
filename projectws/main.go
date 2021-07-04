package main

import (
	"fmt"
	"net/http"
	"projectws/controller"
)

func main() {
	// membuat file server
	dir_name := controller.Dir_Name + "res/"
	fileServer := http.FileServer(http.Dir(dir_name))
	http.Handle("/res/", http.StripPrefix("/res/", fileServer))

	// menghandle path_halaman yang akan divisualisasikan
	http.HandleFunc("/", controller.HomePage)
	http.HandleFunc("/siswa/", controller.PageSiswaShowAll)
	http.HandleFunc("/siswa/insert", controller.PageSiswaInsert)
	http.HandleFunc("/siswa/edit", controller.PageSiswaEdit)
	http.HandleFunc("/siswa/delete", controller.PageSiswaDelete)

	//menjalankan server dengan domain/ip dengan port tertentu
	fmt.Println("Starting webserver on port 8081...")
	http.ListenAndServe("127.0.0.1:8081", nil)
}
