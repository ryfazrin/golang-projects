package pustakasaya


type Siswa struct {
	Nama   string
	Nis    int
	Berat  float32
	Tinggi float32
}

func (s *Siswa) SetTinggi(t float32)  {
	s.Tinggi = t
}
