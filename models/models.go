package models

type Profile struct {
	Nama              string `json:"nama"`
	Alamat            string `json:"alamat"`
	Pekerjaan         string `json:"pekerjaan"`
	AlasanPilihGolang string `json:"alasan"`
}
