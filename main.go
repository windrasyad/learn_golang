package main

import (
	"fmt"
	"os"
	siswa "sesi_4/controllers"
)

func main() {

	args := os.Args[1:2]
	fmt.Println(siswa.BiodataPrint(args))
}
