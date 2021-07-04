package main

import (
	"fmt"
	"net/http"
	"try_db_2/controller"
)

func main() {
	// file server
	dir_name := controller.Dir_Name + "res/"
	fileserver := http.FileServer(http.Dir(dir_name))
	http.Handle("/res/", http.StripPrefix("/res/", fileserver))

	// menghandle halaman utama
	http.HandleFunc("/", controller.Show)

	// menjalankan server
	fmt.Println("server sedang berjalan...")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
