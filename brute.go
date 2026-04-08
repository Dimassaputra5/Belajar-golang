package main

import (
	"bruteforce/function"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Menjalankan Brute Force 4 karakter ASCII Extend(8 bit)")
	target := function.RandomPassword()
	fmt.Printf("Password yang akan di brute force adalah: %s\n", string(target[:]))
	waktuMulai := time.Now()
	result, count := function.Force(target[:])
	waktuEksekusi := time.Since(waktuMulai)
	fmt.Printf("Hasil brute force: %s dengan %s percobaan\n", result, function.ThousandFormat(count))

	fmt.Printf("Waktu yang dibutuhkan: %s\n", waktuEksekusi)
}
