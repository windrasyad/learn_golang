package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"
)

var path = "status.json"

type Status struct {
	Status struct {
		Water int `json:"water"`
		Wind  int `json:"wind"`
	} `json:"status"`
}

type Result struct {
	WaterVol    int    `json:"water_volume"`
	WindSpeed   int    `json:"wind_speed"`
	StatusWater string `json:"status_water"`
	StatusWind  string `json:"status_wind"`
}

func init() {
	go writeFile()
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func writeFile() {
	var _, err = os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> file berhasil dibuat", path)

	// buka file dengan level akses READ & WRITE
	for {
		var file, err = os.OpenFile(path, os.O_RDWR, 0644)
		if err != nil {
			return
		}
		defer file.Close()
		// rand.Seed(time.Now().UnixNano())
		insertFile := Status{}
		insertFile.Status.Water = randomInt(1, 20)
		insertFile.Status.Wind = randomInt(1, 20)
		jsonData, _ := json.Marshal(insertFile)
		// tulis data ke file
		_, err = file.Write(jsonData)
		if err != nil {
			return
		}

		// simpan perubahan
		err = file.Sync()
		if err != nil {
			return
		}

		fmt.Println("==> file berhasil di isi")
		time.Sleep(time.Second * 15)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	// baca file
	result, err := ioutil.ReadAll(file)

	// var data Data
	data := Status{}

	_ = json.Unmarshal([]byte(result), &data)

	var statusWater string
	var statusWind string
	// fmt.Println(data.Status.Water)
	if data.Status.Water <= 5 {
		statusWater = "Aman"
	} else if data.Status.Water >= 6 && data.Status.Water <= 8 {
		statusWater = "Siaga"
	} else if data.Status.Water > 8 {
		statusWater = "Bahaya"
	}

	if data.Status.Wind <= 6 {
		statusWind = "Aman"
	} else if data.Status.Wind >= 7 && data.Status.Wind <= 15 {
		statusWind = "Siaga"
	} else if data.Status.Wind > 15 {
		statusWind = "Bahaya"
	}

	res := Result{
		WaterVol:    data.Status.Water,
		WindSpeed:   data.Status.Wind,
		StatusWater: statusWater,
		StatusWind:  statusWind,
	}
	// tmplt, err := template.New("index").Parse(" Water Volume: {{ .WaterVol }} m \n Wind Speed: {{ .WindSpeed }} m/s \n Water: {{.StatusWater}}, Wind: {{.StatusWind}}")
	tmplt, err := template.ParseFiles("./views/view.gtpl")
	if err != nil {
		log.Fatal(err)
	}
	err = tmplt.Execute(w, res)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	fmt.Println("Listen to addres :9000")
	http.HandleFunc("/", index)
	http.ListenAndServe(":9000", nil)

}
