package main

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
