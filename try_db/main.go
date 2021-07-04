package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type karya struct {
	id    string
	judul string
	isi   string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@/go_karya")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, judul, isi from karya")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []karya

	for rows.Next() {
		var each = karya{}
		var err = rows.Scan(&each.id, &each.judul, &each.isi)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		// fmt.Println(each.id)
		// fmt.Println(each.judul)
		fmt.Printf("nomor karya : %s\n", each.id)
		fmt.Printf("dengan judul : %s\n", each.judul)
	}
}

func main() {
	sqlQuery()
}
