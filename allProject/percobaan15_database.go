package main

import (
	"allProject/pustakasaya"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/sekolah")
	if err != nil {
		panic(0)
	}
	defer db.Close()

	//db.Query("INSERT INTO siswa VALUES ( 2, 'Andi', 710, 80, 157 )")

	s := pustakasaya.Siswa{}
	db.QueryRow("SELECT nama, nis, berat, tinggi FROM siswa").Scan(&s.Nama, &s.Nis, &s.Berat, &s.Tinggi)
	fmt.Println(s)

}
