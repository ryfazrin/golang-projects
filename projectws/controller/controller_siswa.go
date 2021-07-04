package controller

import (
	"fmt"
	"net/http"
	"strconv"
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
	data["siswa"] = ReadDataSiswa()

	RenderTemplate(w, Dir_Name+"siswa_list.html", data)
}

// halaman tambah record siswa
func PageSiswaInsert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		nama := r.PostFormValue("nama")
		nis := r.PostFormValue("nis")
		berat, _ := strconv.ParseFloat(r.PostFormValue("berat"), 64)
		tinggi, _ := strconv.ParseFloat(r.PostFormValue("tinggi"), 64)
		fmt.Println(nama, nis, berat, tinggi)
		CreateDataSiswa("", nama, nis, berat, tinggi)

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
		nama := r.PostFormValue("nama")
		nis := r.PostFormValue("nis")
		berat, _ := strconv.ParseFloat(r.PostFormValue("berat"), 64)
		tinggi, _ := strconv.ParseFloat(r.PostFormValue("tinggi"), 64)
		fmt.Println(id, nama, nis, berat, tinggi)
		UpdateDataSiswaById(id, nama, nis, berat, tinggi)
		data["info"] = "Data berhasil diupdate"
	}
	data["siswa"] = ReadDataSiswaById(id)
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
