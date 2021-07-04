package controller

type Karya struct {
	id    string
	judul string
	isi   string
}

// create or insert a record to table siswa
func CreateDataSiswa(id string, judul string, isi string) {
	db := Conn()
	defer db.Close()
	stmt, _ := db.Prepare("INSERT INTO karya VALUES ( ?, ?, ?, ?, ? )")
	stmt.Exec(id, judul, isi)
	defer stmt.Close()
}

// read all data siswa
func ReadDataSiswa() []Karya {
	db := Conn()
	defer db.Close()
	rows, err := db.Query("SELECT id, judul, isi FROM karya")
	if err != nil {
		panic(err.Error())
	}

	all_siswa := []Karya{}
	for rows.Next() {
		s := Karya{}
		err = rows.Scan(&s.id, &s.judul, &s.isi)
		if err != nil {
			panic(err.Error())
		}
		all_siswa = append(all_siswa, s)
	}
	return all_siswa
}

// read only one row
func ReadDataSiswaById(id string) Karya {
	db := Conn()
	defer db.Close()
	s := Karya{}
	db.QueryRow("SELECT id, judul, isi FROM karya WHERE id=?", id).Scan(&s.id, &s.judul, &s.isi)
	return s
}

// update record siswa by id
func UpdateDataSiswaById(id string, judul string, isi string) {
	db := Conn()
	defer db.Close()
	stmt, _ := db.Prepare("UPDATE karya SET judul=?, isi=? WHERE id=?")
	stmt.Exec(judul, isi, id)
	defer stmt.Close()
}

// remove data siswa by id
func DeleteDataSiswa(id string) {
	db := Conn()
	defer db.Close()
	stmt, _ := db.Prepare("DELETE FROM karya WHERE id=?")
	stmt.Exec(id)
	defer stmt.Close()
}
