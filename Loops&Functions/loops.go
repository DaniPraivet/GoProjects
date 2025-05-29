package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	x := float64(43)
	fmt.Println("Square root of", x, ":", Sqrt(x))
}
