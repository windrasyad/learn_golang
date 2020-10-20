package models

type Profile struct {
	Nama      string `json:"nama"`
	Umur      int    `json:"umur"`
	Alamat    Alamat `json:"alamat"`
	Pekerjaan string `json:"pekerjaan"`
	Agama     string `json:"agama"`
}

type Alamat struct {
	NamaJalan string `json:"namaJalan"`
	Kelurahan string `json:"kelurahan"`
	Kecamatan string `json:"kecamatan"`
	RT        string `json:"RT"`
	RW        string `json:"RW"`
	Kota      string `json:"kota"`
	Provinsi  string `json:"provinsi"`
	KodePos   int32  `json:"kodePos"`
}
