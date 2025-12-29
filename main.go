package main

import (
	"fmt"
	"math/rand"
	"time"
)

// GCD Iteratif
func gcdIteratif(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// GCD Rekursif
func gcdRekursif(a, b int) int {
	if b == 0 {
		return a
	}
	return gcdRekursif(b, a%b)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== Analisis Efisiensi Otomatis ===")
	inputSizes := []int{1, 10, 100, 1000, 5000, 10000}
	const repeat = 1000000 // dinaikkan dari 100.000 ke 1 juta

	// Header tabel rapi
	fmt.Printf("%-8s %-15s %-15s\n", "N", "Iteratif (ns)", "Rekursif (ns)")

	for _, N := range inputSizes {
		// Pilih bilangan acak dari 1 sampai N
		x := rand.Intn(N) + 1
		y := rand.Intn(N) + 1

		// Iteratif
		start := time.Now()
		for i := 0; i < repeat; i++ {
			_ = gcdIteratif(x, y)
		}
		iterTime := time.Since(start).Nanoseconds()

		// Rekursif
		start = time.Now()
		for i := 0; i < repeat; i++ {
			_ = gcdRekursif(x, y)
		}
		recTime := time.Since(start).Nanoseconds()

		// Cetak hasil rapi
		fmt.Printf("%-8d %-15d %-15d\n", N, iterTime, recTime)
	}

	fmt.Println("\nSelesai! Data siap untuk dibuat grafik.")
}
