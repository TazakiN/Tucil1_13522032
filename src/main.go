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

func isUnique(buf []coor) bool {
	//cek apakah setiap koordinat di buf unik
	for i := 0; i < len(buf); i++ {
		for j := i + 1; j < len(buf); j++ {
			if buf[i] == buf[j] {
				return false
			}
		}
	}
	return true
}

func isValid(buf []coor) bool {
	// cek apakah buf valid
	for i := 1; i < len(buf); i++ {
		if i%2 == 1 {
			if buf[i].Y != buf[i-1].Y {
				return false
			}
		} else {
			if buf[i].X != buf[i-1].X {
				return false
			}
		}
	}
	return true
}

func compareBuffer(buffer1 []coor, buffer2 []coor) bool {
	//bandingkan bufferTemp dengan bufferTertinggi
	for i := 0; i < len(bufferTemp); i++ {
		if bufferTemp[i] != bufferTertinggi[i] {
			return false
		}
	}
	return true
}

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
			fmt.Printf("%d, %d\n", bufferTertinggi[i].X, bufferTertinggi[i].Y)
		}

	} else {
		fmt.Println("\033[31mTidak ada buffer yang valid\033[0m")
	}
	fmt.Print("\033[35mWaktu eksekusi: \033[0m")
	fmt.Println(elapsedTime)
}
