package controller

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Karya struct {
	id    string
	judul string
	isi   string
}

var Dir_Name = "C:/GOPATH/src/try_db_2/"
var BaseURL = "http://localhost:8080/"

func Conn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@/go_karya")
	if err != nil {
		return nil, err
	}

	return db, nil
}

// File Reader (output dalam bentuk []byte)
func ReadHtmlFile(name string) []byte {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		panic(0)
	}
	return data
}

func ReadDataKarya() []Karya {
	db, err := Conn()
	defer db.Close()
	rows, err := db.Query("SELECT id, judul, isi FROM karya")
	if err != nil {
		panic(err.Error())
	}

	all_karya := []Karya{}
	for rows.Next() {
		s := Karya{}
		err = rows.Scan(&s.id, &s.judul, &s.isi)
		if err != nil {
			panic(err.Error())
		}
		all_karya = append(all_karya, s)
	}
	return all_karya
}

// // read all data karya
// func ReadDataKarya() []Karya {
// 	db, err := Conn()
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT * FROM karya")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer rows.Close()

// 	all_karya := []Karya{}

// 	for rows.Next() {
// 		k := Karya{}
// 		err = rows.Scan(&k.id, &k.judul, &k.isi)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		all_karya = append(all_karya, k)
// 	}
// 	return all_karya

// 	// if err = rows.Err(); err != nil {
// 	// 	fmt.Println(err.Error())
// 	// 	return
// 	// }

// 	// for _, k := range all_karya {
// 	// 	// fmt.Println(each.id)
// 	// 	// fmt.Println(each.judul)
// 	// 	fmt.Printf("nomor karya : %s\n", k.id)
// 	// 	fmt.Printf("dengan judul : %s\n", k.judul)
// 	// }
// }
