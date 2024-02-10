package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func displayHasil(elapsedTime time.Duration) {
	// tampilkan hasil
	if bufferTertinggiReward != 0 {
		fmt.Print("\033[35m\nReward Tertinggi: \033[0m")
		fmt.Println(bufferTertinggiReward)
		fmt.Print("\033[35mToken Buffer Terbaik: \033[0m")
		for i := 0; i < len(bufferTertinggi); i++ {
			fmt.Print(mat[bufferTertinggi[i].X][bufferTertinggi[i].Y], " ")
		}
		fmt.Println("\033[35\nmBuffer Terbaik: \033[0m")
		for i := 0; i < len(bufferTertinggi); i++ {
			fmt.Printf("%d, %d\n", bufferTertinggi[i].X+1, bufferTertinggi[i].Y+1)
		}

	} else {
		fmt.Println("\033[31mTidak ada buffer yang valid\033[0m")
	}
	fmt.Print("\033[35mWaktu eksekusi: \033[0m")
	fmt.Println(elapsedTime)

	// tawarkan untuk menyimpan hasil ke file txt
	fmt.Print("\033[34mApakah Anda ingin menyimpan hasil ke file? (y/n): \033[0m")
	var input string
	fmt.Scanln(&input)
	if input == "y" {
		saveToFile()
	}
	fmt.Println("\033[31mTerima kasih telah menggunakan program ini\033[0m")
}

func saveToFile() {

	// minta user untuk memasukkan nama file .txt
	fmt.Print("\033[34mMasukkan nama file .txt (tanpa .txt): \033[0m")
	var input string
	fmt.Scanln(&input)

	// validasi keberadaan file
	_, err := os.Stat("../test/" + input + ".txt")
	if err == nil {
		fmt.Println("\033[31mFile sudah ada\033[0m")
		saveToFile()
	}
	// buat file
	file, err := os.Create("../test/" + input + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// tulis hasil ke file
	file.WriteString("Reward Tertinggi: " + strconv.Itoa(bufferTertinggiReward) + "\n")
	file.WriteString("Token Buffer Terbaik: ")
	for i := 0; i < len(bufferTertinggi); i++ {
		file.WriteString(mat[bufferTertinggi[i].X][bufferTertinggi[i].Y] + " ")
	}
	file.WriteString("\nBuffer Terbaik: \n")
	for i := 0; i < len(bufferTertinggi); i++ {
		file.WriteString(strconv.Itoa(bufferTertinggi[i].X+1) + ", " + strconv.Itoa(bufferTertinggi[i].Y+1) + "\n")
	}
	file.WriteString("Waktu eksekusi: " + elapsedTime.String())
	fmt.Println("\033[33mMenyimpan file...\033[0m")

}
