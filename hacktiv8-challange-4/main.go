package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
}

var temanKelas = []Teman{
	{"Syarif Fadilla", "Mertapadawetan", "IT"},
	{"Usman Habib", "Blok Pon2", "Back-end"},
	{"Ajis", "Japura", "UI/UX"},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run biodata.go [nomor absen]")
		os.Exit(1)
	}

	nomorAbsen := os.Args[1]

	teman, err := getTeman(nomorAbsen)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
}

func getTeman(nomorAbsen string) (Teman, error) {
	var teman Teman

	absen, err := strconv.Atoi(nomorAbsen)
	if err != nil || absen < 1 || absen > len(temanKelas) {
		return teman, fmt.Errorf("Nomor absen tidak valid")
	}

	return temanKelas[absen-1], nil
}
