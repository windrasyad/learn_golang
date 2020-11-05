package controllers

import (
	"encoding/json"
	"sesi_4/models"
)

func BiodataPrint(id []string) interface{} {
	if id[0] == "1" {
		profile := models.Profile{
			Nama:              "Rahmat Faisal",
			Alamat:            "Jakarta Pusat",
			Pekerjaan:         "Golang",
			AlasanPilihGolang: "Menyenangkan",
		}
		a, _ := json.Marshal(profile)
		return string(a)

	} else if id[0] == "2" {
		profile := models.Profile{
			Nama:              "Irma Nur Afifah",
			Alamat:            "Njagir",
			Pekerjaan:         "Presiden",
			AlasanPilihGolang: "Menantang",
		}

		a, _ := json.Marshal(profile)
		return string(a)

	} else if id[0] == "3" {
		profile := models.Profile{
			Nama:              "Windra Rasyad",
			Alamat:            "Jakarta Selatan",
			Pekerjaan:         "Software Developers",
			AlasanPilihGolang: "Innovative",
		}

		a, _ := json.Marshal(profile)
		return string(a)

	} else if id[0] == "4" {
		profile := models.Profile{
			Nama:              "Andre Philo",
			Alamat:            "Jakarta Pusat",
			Pekerjaan:         "Golang",
			AlasanPilihGolang: "Menantang",
		}

		a, _ := json.Marshal(profile)
		return string(a)

	} else if id[0] == "5" {
		profile := models.Profile{
			Nama:              "Reydika Ilham",
			Alamat:            "Malang",
			Pekerjaan:         "Kepala Divisi",
			AlasanPilihGolang: "Menantang",
		}

		a, _ := json.Marshal(profile)
		return string(a)

	} else {
		return "Tidak Ada Siswa dengan Nomor ID Tersebut"
	}
}
