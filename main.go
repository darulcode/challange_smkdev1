package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	//input dari pengguna
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, err := strconv.Atoi(scanner.Text())

	// mengecek inputan pengguna benar adalah integer
	if err != nil {
		panic(err)
	}

	//perulangan untuk menampilkan hasil dari dadu dengan user input kali perulangan
	for i := 1; i <= input; i++ {
		//mendapatkan hasil cube
		hasilCube := 0
		for j := 1; j <= i; j++ {
			hasilCube += cube()
		}

		// menampilkan hasil
		fmt.Printf("Current Number is : %d and the cube is %d\n", i, hasilCube)
	}
}

func cube() int {
	// Inisialisasi generator angka acak dengan seed berdasarkan waktu
	rand.Seed(time.Now().UnixNano())

	// Menggunakan rand.Intn untuk mendapatkan angka acak antara 1 dan 6
	randomNumber := rand.Intn(6) + 1

	return randomNumber
}
