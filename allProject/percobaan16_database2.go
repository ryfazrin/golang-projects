package main

import (
	"allProject/pustakasaya"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// contoh 1
	CreateDataSiswa("2", "Andi", "702", 56, 58)
	UpdateDataSiswaById(2, "Andi", "705", 56, 56)
	//DeleteDataSiswa("2")

	//contoh 2
	data := ReadDataSiswa()
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i].Nama, data[i].Nis, data[i].Berat, data[i].Tinggi)
	}

	// contoh 3
	s := ReadDataSiswaById(1)
	fmt.Println("\nGet row by id siswa = 1:", s.Nama, s.Nis, s.Berat, s.Tinggi)
}

// create connection
func Conn() *sql.DB {
	db, err := sql.Open("mysql", "root:@/sekolah")
	if err != nil {
		panic(0)
	}
	return db
}

// create or insert a record to table siswa
func CreateDataSiswa(id string, name string, nis string, berat float32, tinggi float32) {
	db := Conn()
	defer db.Close()
	stmt, _ := db.Prepare("INSERT INTO siswa VALUES ( ?, ?, ?, ?, ? )")
	stmt.Exec(id, name, nis, berat, tinggi)
	defer stmt.Close()
}

// read all data siswa
func ReadDataSiswa() []pustakasaya.Siswa {
	db := Conn()
	defer db.Close()
	rows, err := db.Query("SELECT nama, nis, berat, tinggi FROM siswa")
	if err != nil {
		panic(err.Error())
	}

	all_siswa := []pustakasaya.Siswa{}
	for rows.Next() {
		s := pustakasaya.Siswa{}
		err = rows.Scan(&s.Nama, &s.Nis, &s.Berat, &s.Tinggi)
		if err != nil {
			panic(err.Error())
		}
		all_siswa = append(all_siswa, s)
	}
	return all_siswa
}

// read only one row
func ReadDataSiswaById(id int) pustakasaya.Siswa {
	db := Conn()
	defer db.Close()
	s := pustakasaya.Siswa{}
	db.QueryRow("SELECT nama, nis, berat, tinggi FROM siswa WHERE id=?", id).Scan(&s.Nama, &s.Nis, &s.Berat, &s.Tinggi)
	return s
}

// update record siswa by id
func UpdateDataSiswaById(id int, name string, nis string, berat float32, tinggi float32) {
	db := Conn()
	defer db.Close()
	stmt, _ := db.Prepare("UPDATE siswa SET nama=?, nis=?, berat=?, tinggi=? WHERE id=?")
	stmt.Exec(name, nis, berat, tinggi, id)
	defer stmt.Close()
}

// remove data siswa by id
func DeleteDataSiswa(id string) {
	db := Conn()
	defer db.Close()
	stmt, _ := db.Prepare("DELETE FROM siswa WHERE id=?")
	stmt.Exec(id)
	defer stmt.Close()
}
