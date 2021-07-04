package controller

import (
	"net/http"
)

func Show(w http.ResponseWriter, r *http.Request) {
	// membuat penampung data agar dapat dipanggil melalui template html
	data := make(map[string]interface{})

	// mengambil data dari database
	data["karya"] := ReadDataKarya()

	RenderTemplate(w, Dir_Name+"index.html", data)

}
