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
	elapsedTime           time.Duration
)

func main() {
	printAsciiArt()
	mainInput()

	// inisisasi variabel
	bufferTertinggi = make([]coor, bufferSize)
	bufferTemp = make([]coor, bufferSize)
	bufferStopper := make([]coor, bufferSize)
	bufferTertinggiReward = 0
	bufferTempReward = 0
	stopper = 0

	fmt.Println("\033[33mMemulai kalkulasi...\033[0m")
	startTime := time.Now()
	for stopper < 1 {
		// cek apakah bufferTemp valid
		if isUnique(bufferTemp) && isValid(bufferTemp) {
			bufferTempReward = hitungReward(bufferTemp)
			if bufferTempReward > bufferTertinggiReward {
				bufferTertinggiReward = bufferTempReward
				copy(bufferTertinggi, bufferTemp)
			}
		}
		// fmt.Println("bufferTemp: ", bufferTemp)
		bufferTemp = nextBufferTemp(bufferTemp)
		if compareBuffer(bufferTemp, bufferStopper) {
			stopper++
		}
	}

	elapsedTime = time.Since(startTime)
	displayHasil(elapsedTime)
	fmt.Println("\033[31mKeluar dari program...\033[0m")
}
