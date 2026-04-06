package function

import "math/rand"

func Force(pass []byte) (string, int) {
	p := []byte(pass)
	iter := 0
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			for k := 0; k < 256; k++ {
				for l := 0; l < 256; l++ {
					iter++
					if byte(i) == p[0] && byte(j) == p[1] && byte(k) == p[2] && byte(l) == p[3] {
						return string([]byte{byte(i), byte(j), byte(k), byte(l)}), iter
					}

				}
			}
		}
	}
	return "Not Found", iter
}

func RandomPassword() [4]byte {
	var password [4]byte
	for i := 0; i < 4; i++ {
		password[i] = byte(rand.Intn(256))
	}
	return password
}
