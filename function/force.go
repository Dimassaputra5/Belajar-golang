package function

import (
	"math/rand"
	"strconv"
)

func Force(pass []byte) (string, int) {
	p := []byte(pass)
	iter := 0
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			for k := 0; k < 256; k++ {
				for l := 0; l < 256; l++ {
					iter++
					if byte(i) == p[0] && byte(j) == p[1] &&
						byte(k) == p[2] && byte(l) == p[3] {
						return string([]byte{byte(i), byte(j), byte(k), byte(l)}), iter
					}

				}
			}
		}
	}
	return "Not Found", iter
}

func RandomPassword(n string) []byte {
	p, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	var password []byte
	for i := 0; i < p; i++ {
		password = append(password, byte(rand.Intn(256)))
	}
	return password
}

func ThousandFormat(num int) string {
	str := strconv.Itoa(num)
	length := len(str)
	if length <= 3 {
		return str
	}
	var result string
	for i, digit := range str {
		if (length-i)%3 == 0 && i != 0 {
			result += ","
		}
		result += string(digit)
	}
	return result
}

func Force5(pass []byte) (string, int) {
	p := []byte(pass)
	iter := 0
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			for k := 0; k < 256; k++ {
				for l := 0; l < 256; l++ {
					for m := 0; m < 256; m++ {
						iter++
						if byte(i) == p[0] && byte(j) == p[1] &&
							byte(k) == p[2] && byte(l) == p[3] && byte(m) == p[4] {
							return string([]byte{byte(i), byte(j), byte(k), byte(l), byte(m)}), iter
						}

					}
				}
			}
		}
	}
	return "Not Found", iter
}
