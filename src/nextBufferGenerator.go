package main

func nextCoor(koor coor, horizontal bool) coor {

	if horizontal {
		if koor.X+1 < matrixWidth {
			koor.X++
		} else {
			return coor{-1, -1}
		}
	} else {
		if koor.Y+1 < matrixHeight {
			koor.Y++
		} else {
			return coor{-1, -1}
		}
	}

	return koor
}

func nextBufferTemp(bufferTemp []coor) []coor {
	// generate bufferTemp selanjutnya
	// validasi bufferTemp
	for j := 0; j < len(bufferTemp)-1; j++ {
		if bufferTemp[j].X != bufferTemp[j+1].X && j%2 == 1 {
			bufferTemp[j+1] = coor{bufferTemp[j].X, bufferTemp[j+1].Y}
			return bufferTemp
		} else if bufferTemp[j].Y != bufferTemp[j+1].Y && j%2 == 0 {
			bufferTemp[j+1] = coor{bufferTemp[j+1].X, bufferTemp[j].Y}
			return bufferTemp
		}
	}

	// jika bufferTemp tidak valid, generate bufferTemp selanjutnya
	for i := len(bufferTemp) - 1; i >= 0; i-- {
		if i%2 == 1 { // jika i ganjil, maka bufferTemp[i] adalah horizontal
			bufferTemp[i] = nextCoor(bufferTemp[i], true)
			if bufferTemp[i].X == -1 && bufferTemp[i].Y == -1 {
				bufferTemp[i] = coor{0, 0}
			} else {
				return bufferTemp
			}
		} else { // jika i genap, maka bufferTemp[i] adalah vertikal
			bufferTemp[i] = nextCoor(bufferTemp[i], false)
			if bufferTemp[i].X == -1 && bufferTemp[i].Y == -1 {
				bufferTemp[i] = coor{0, 0}
			} else {
				return bufferTemp
			}
		}

	}
	return bufferTemp
}
