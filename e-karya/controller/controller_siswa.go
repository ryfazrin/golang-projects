package controller

import (
	"fmt"
	"net/http"
)

// fungsi untuk menangani halaman utama
func HomePage(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, Dir_Name+"index.html", nil)
}

// halaman list data siswa
func PageSiswaShowAll(w http.ResponseWriter, r *http.Request) {
	// membuat tempat penampungan data siswa supaya dapat dipanggil melalui template html
	data := make(map[string]interface{})

	// mengambil data siswa dari database
	data["karya"] = ReadDataSiswa()

	RenderTemplate(w, Dir_Name+"siswa_list.html", data)
}

// halaman tambah record siswa
func PageSiswaInsert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		judul := r.PostFormValue("judul")
		isi := r.PostFormValue("isi")
		fmt.Println(judul, isi)
		CreateDataSiswa("", judul, isi)

		data := make(map[string]interface{})
		data["info"] = "Data berhasil ditambahkan ke database"
		RenderTemplate(w, Dir_Name+"siswa_insert.html", data)
	} else {
		RenderTemplate(w, Dir_Name+"siswa_insert.html", nil)
	}
}

// halaman edit record siswa
func PageSiswaEdit(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	r.ParseForm()
	id := r.FormValue("id")
	if r.Method == "POST" {
		judul := r.PostFormValue("judul")
		isi := r.PostFormValue("isi")
		fmt.Println(id, judul, isi)
		UpdateDataSiswaById(id, judul, isi)
		data["info"] = "Data berhasil diupdate"
	}
	data["karya"] = ReadDataSiswaById(id)
	RenderTemplate(w, Dir_Name+"siswa_edit.html", data)
}

// halaman hapus record siswa
func PageSiswaDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	DeleteDataSiswa(id)
	data := make(map[string]interface{})
	data["info"] = "Data dengan id=" + id + " berhasil dihapus!"
	RenderTemplate(w, Dir_Name+"siswa_delete.html", data)
}
