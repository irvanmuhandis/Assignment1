package main

import (
	"Assignment1/helpers"
	"fmt"
	"os"
)

func main() {
	// Ambil argumen dari command line
	args := os.Args

	// Pastikan argumen nomor absen disediakan
	if len(args) < 2 {
		fmt.Println("Usage: go run biodata.go [No]")
		return
	}

	// Parsing nomor absen dari argumen
	var no int
	_, err := fmt.Sscanf(args[1], "%d", &no)
	if err != nil {
		fmt.Println("Invalid argument for No")
		return
	}

	// Dapatkan data teman berdasarkan nomor absen
	friend := helpers.GetFriendData(no)
	if friend == nil {
		fmt.Println("No data found for No", no)
		return
	}

	// Tampilkan data teman
	fmt.Println("Nama:", friend.Nama)
	fmt.Println("Alamat:", friend.Alamat)
	fmt.Println("Pekerjaan:", friend.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", friend.Alasan)
}
