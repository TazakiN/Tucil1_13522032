package main

func hitungReward(bufferTemp []coor) int {
	reward := 0
	for i := 0; i < len(bufferTemp); i++ {
		// hitung reward dari bufferTemp
		for j := 0; j < len(seqs); j++ {
			if mat[bufferTemp[i].X][bufferTemp[i].Y] == seqs[j][0] {
				// jika huruf pertama cocok, cek apakah sisa huruf juga cocok
				if cekSeq(i, j) == 1 {
					reward += rewardSeqs[j]
				}
			}
		}
	}
	return reward
}

func cekSeq(buffPos int, seqIdx int) int {
	seqPos := 1
	// cek apakah sisa huruf juga cocok
	for i := buffPos + 1; i < len(bufferTemp); i++ {
		if mat[bufferTemp[i].X][bufferTemp[i].Y] != seqs[seqIdx][seqPos] {
			return 0
		}
		seqPos++
		if seqPos == len(seqs[seqIdx]) {
			return 1
		}
	}
	return 0
}
