package main

import "net/http"

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := []byte("Hello world!!!") // dapat juga ditulis ==> io.WriteString(w, "Hello world!!!")
	w.Write(data)
}
