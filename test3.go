package main

import "fmt"

func main() {
	var a, b, h float64
	a = 10 // sisi sejajar pertama
	b = 20 // sisi sejajar kedua
	h = 5  // tinggi

	luas := 0.5 * (a + b) * h
	fmt.Printf("Luas trapesium dengan sisi sejajar %.2f dan %.2f serta tinggi %.2f adalah %.2f\n", a, b, h, luas)
}
