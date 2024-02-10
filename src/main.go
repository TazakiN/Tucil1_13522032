package main

import (
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

	elapsedTime = time.Since(startTime)
	displayHasil(elapsedTime)
}
