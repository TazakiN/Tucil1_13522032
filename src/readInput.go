package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func readFile(input string) {
	// buka file input.txt
	f, err := os.Open("../test/" + input + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// line pertama adalah bufferSize
	scanner.Scan()
	bufferSize, _ = strconv.Atoi(scanner.Text())

	// line kedua adalah matrixSize, misal '6 6' (6 baris, 6 kolom), buat menjadi 2 variabel
	scanner.Scan()
	matrixWidth, _ = strconv.Atoi(strings.Split(scanner.Text(), " ")[0])
	matrixHeight, _ = strconv.Atoi(strings.Split(scanner.Text(), " ")[1])

	// line selanjutnya adalah matrix, lakukan proses pembacaan dan simpan dalam variabel mat
	for i := 0; i < matrixHeight; i++ {
		scanner.Scan()
		row := strings.Split(scanner.Text(), " ")
		mat = append(mat, row)
	}

	scanner.Scan()
	banyakSeq, _ = strconv.Atoi(strings.Split(scanner.Text(), " ")[0])
	seqs = make([][]string, banyakSeq)
	rewardSeqs = make([]int, banyakSeq)

	for i := 0; i < banyakSeq; i++ {
		scanner.Scan()
		seqs[i] = strings.Split(scanner.Text(), " ")
		scanner.Scan()
		rewardSeqs[i], _ = strconv.Atoi(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func readManualInput() {
	fmt.Print("\033[34mMasukkan banyak token unik: \033[0m")
	banyakTokenUnik := 0
	fmt.Scanln(&banyakTokenUnik)

	fmt.Print("\033[34mMasukkan token-token unik: \033[0m")
	tokens := make([]string, banyakTokenUnik)
	for i := 0; i < banyakTokenUnik; i++ {
		fmt.Scan(&tokens[i])
	}
	fmt.Scanln() // untuk mengakhiri newline

	fmt.Print("\033[34mMasukkan Ukuran Buffer: \033[0m")
	fmt.Scanln(&bufferSize)

	fmt.Print("\033[34mMasukkan Ukuran Matrix (baris kolom): \033[0m")
	fmt.Scanln(&matrixHeight, &matrixWidth)

	fmt.Print("\033[34mMasukkan banyak sekuens: \033[0m")
	fmt.Scanln(&banyakSeq)

	fmt.Print("\033[34mMasukkan panjang maksimal sekuens: \033[0m")
	panjangMaksimalSekuens := 0
	fmt.Scanln(&panjangMaksimalSekuens)

	// bentuk matrix secara acak dengan setiap elemen adalah salah satu token unik
	mat = make([][]string, matrixHeight)
	for i := 0; i < matrixHeight; i++ {
		mat[i] = make([]string, matrixWidth)
		for j := 0; j < matrixWidth; j++ {
			mat[i][j] = tokens[rand.Intn(banyakTokenUnik)]
		}
	}

	// bentuk sekuens secara acak dengan panjang maksimal sekuens yang dimasukkan
	seqs = make([][]string, banyakSeq)
	rewardSeqs = make([]int, banyakSeq)
	for i := 0; i < banyakSeq; i++ {
		seqLength := rand.Intn(panjangMaksimalSekuens-1) + 2
		seqs[i] = make([]string, seqLength)
		for j := 0; j < seqLength; j++ {
			seqs[i][j] = tokens[rand.Intn(banyakTokenUnik)]
		}
		rewardSeqs[i] = rand.Intn(100)
	}

	// cetak matrix ke layar
	fmt.Println("\033[35mMatrix: \033[0m")
	for i := 0; i < matrixHeight; i++ {
		for j := 0; j < matrixWidth; j++ {
			fmt.Print(mat[i][j], " ")
		}
		fmt.Println()
	}

	// cetak sekuens ke layar
	fmt.Println("\033[35mSekuens: \033[0m")
	for i := 0; i < banyakSeq; i++ {
		fmt.Print("\033[32mSekuens ", i+1, ": \033[0m")
		for j := 0; j < len(seqs[i]); j++ {
			fmt.Print(seqs[i][j], " ")
		}
		fmt.Println("Reward:", rewardSeqs[i])
	}
}

func printAsciiArt() {
	// Open the file
	file, err := os.Open("ascii.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func mainInput() {

	printAsciiArt()

	// tanya kepada user apakah ingin menggunakan file input.txt atau tidak
	fmt.Println("\033[34mApakah anda ingin menggunakan file input.txt? (y/n)\033[0m")
	var input string
	fmt.Scanln(&input)
	if input == "y" {
		// minta user untuk memasukkan nama file .txt
		fmt.Println("\033[34mMasukkan nama file .txt (tanpa .txt)\033[0m")
		fmt.Scanln(&input)
		// validasi keberadaan file
		_, err := os.Stat("../test/" + input + ".txt")
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("\033[31mFile tidak ditemukan\033[0m")
				mainInput()
			}
		}
		fmt.Println("\033[33mMembaca file...\033[0m")
		readFile(input)
	} else {
		fmt.Println("\033[33mberalih ke penggunaan input manual...\033[0m")
		readManualInput()
	}
}
