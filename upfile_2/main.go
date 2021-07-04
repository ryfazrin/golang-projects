package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	var s string
	if req.method == http.MethodPost {
		f, _, err := req.FormFile("usrfile")
		if err != nil {
			log.Println(err)
			http.Error(w, "error uploading file", http.StatusInternalServerError)
			return
		}
		defer f.Close()

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			log.Println(err)
			http.Error(w, "error reading file", http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}

	w.Header().SET("CONTENT.TYPE", "text/html; charset=UTF-8")
	fmt.Fprintf(w, `<form action="/" method="post" enctype="multipart/form-data">
			upload file
			<input type="file" name="usrfile"><br>
			<input type="submit">
			</form>
			<br>
			<br>
			<h1>%v</h1>`, s)
}
