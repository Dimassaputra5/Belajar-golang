package main

import (
	"bruteforce/function"
	"fmt"
	"time"
)

func template(n string, forceFunc func([]byte) (string, int)) {
	fmt.Printf("Menjalankan Brute Force %s karakter ASCII Extend(8 bit)\n", n)
	target := function.RandomPassword(n)
	fmt.Printf("Password yang akan di brute force adalah: %s\n", string(target[:]))
	waktuMulai := time.Now()
	result, count := forceFunc(target[:])
	waktuEksekusi := time.Since(waktuMulai)
	fmt.Printf("Hasil brute force: %s dengan %s percobaan\n", result, function.ThousandFormat(count))

	fmt.Printf("Waktu yang dibutuhkan: %s\n", waktuEksekusi)
}

func main() {
	template("4", function.Force)
	// template("5", function.Force5)
}
