package main

import (
	"bruteforce/function"
	"fmt"
)

func main() {
	fmt.Println("Menjalankan Brute Force 4 karakter ASCII Extend(8 bit)")
	target := function.RandomPassword()
	fmt.Printf("Password yang akan di brute force adalah: %s\n", string(target[:]))
	result, count := function.Force(target[:])
	fmt.Printf("Hasil brute force: %s dengan %s percobaan\n", result, function.HundredFormat(count))
}
