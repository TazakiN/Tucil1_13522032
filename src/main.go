package main

import (
	"fmt"
	"time"
)

type coor struct {
	X, Y int
}

// KAMUS GLOBAL
var (
	bufferSize            int
	matrixWidth           int
	matrixHeight          int
	mat                   [][]string
	seqs                  [][]string
	rewardSeqs            []int
	banyakSeq             int
	bufferTertinggi       []coor
	bufferTertinggiReward int
	bufferTemp            []coor
	bufferTempReward      int
	stopper               int
)

func main() {
	mainInput()

	// inisisasi variabel
	bufferTertinggi = make([]coor, bufferSize)
	bufferTemp = make([]coor, bufferSize)
	bufferStopper := make([]coor, bufferSize)
	bufferTertinggiReward = 0
	bufferTempReward = 0

	startTime := time.Now()
	for stopper != 2 {
		// cek apakah bufferTemp valid
		if isUnique(bufferTemp) && isValid(bufferTemp) {
			bufferTempReward = hitungReward(bufferTemp)
			if bufferTempReward > bufferTertinggiReward {
				bufferTertinggiReward = bufferTempReward
				copy(bufferTertinggi, bufferTemp)
			}
		}
		bufferTemp = nextBufferTemp(bufferTemp)
		if compareBuffer(bufferTemp, bufferStopper) {
			stopper++
		}
	}

	elapsedTime := time.Since(startTime)

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
}
