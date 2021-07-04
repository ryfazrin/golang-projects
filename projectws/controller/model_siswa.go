package controller

type Siswa struct {
	Id     string
	Nama   string
	Nis    int
	Berat  float64
	Tinggi float64
}

// create or insert a record to table siswa
func CreateDataSiswa(id string, name string, nis string, berat float64, tinggi float64) {
	db := Conn()
	defer db.Close()
	stmt, _ := db.Prepare("INSERT INTO siswa VALUES ( ?, ?, ?, ?, ? )")
	stmt.Exec(id, name, nis, berat, tinggi)
	defer stmt.Close()
}

// read all data siswa
func ReadDataSiswa() []Siswa {
	db := Conn()
	defer db.Close()
	rows, err := db.Query("SELECT id, nama, nis, berat, tinggi FROM siswa")
	if err != nil {
		panic(err.Error())
	}

	all_siswa := []Siswa{}
	for rows.Next() {
		s := Siswa{}
		err = rows.Scan(&s.Id, &s.Nama, &s.Nis, &s.Berat, &s.Tinggi)
		if err != nil {
			panic(err.Error())
		}
		all_siswa = append(all_siswa, s)
	}
	return all_siswa
}

// read only one row
func ReadDataSiswaById(id string) Siswa {
	db := Conn()
	defer db.Close()
	s := Siswa{}
	db.QueryRow("SELECT id, nama, nis, berat, tinggi FROM siswa WHERE id=?", id).Scan(&s.Id, &s.Nama, &s.Nis, &s.Berat, &s.Tinggi)
	return s
}

// update record siswa by id
func UpdateDataSiswaById(id string, name string, nis string, berat float64, tinggi float64) {
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
