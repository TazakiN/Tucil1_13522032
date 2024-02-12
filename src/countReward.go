package main

func hitungReward(bufferTemp []coor) int {
	reward := 0
	for i := 0; i < len(bufferTemp); i++ {
		// hitung reward dari bufferTemp
		for j := 0; j < len(seqs); j++ {
			if mat[bufferTemp[i].Y][bufferTemp[i].X] == seqs[j][0] {
				// jika huruf pertama cocok, cek apakah sisa huruf juga cocok
				if cekSeq(i, j) {
					reward += rewardSeqs[j]
				}
			}
		}
	}
	return reward
}

func cekSeq(buffPos int, seqIdx int) bool {
	seqPos := 1
	// cek apakah sisa huruf juga cocok
	for i := buffPos + 1; i < len(bufferTemp); i++ {
		if mat[bufferTemp[i].Y][bufferTemp[i].X] != seqs[seqIdx][seqPos] {
			return false
		}
		seqPos++
		if seqPos == len(seqs[seqIdx]) {
			return true
		}
	}
	return false
}
