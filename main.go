package main

import (
	"encoding/json"
	"fmt"
	"sesi_1/models"
)

func main() {

	profil := models.Profile{
		Nama: "Windra Rasyad",
		Umur: 23,
		Alamat: models.Alamat{
			NamaJalan: "H.Sinen",
			Kelurahan: "Ragunan",
			Kecamatan: "Pasar Minggu",
			RT:        "01",
			RW:        "06",
			Kota:      "Jakarta Selatan",
			Provinsi:  "DKI Jakarta",
			KodePos:   12550,
		},
		Pekerjaan: "Karyawan",
		Agama:     "Islam",
	}

	profilAll, _ := json.Marshal(profil)
	fmt.Println(string(profilAll))

}
