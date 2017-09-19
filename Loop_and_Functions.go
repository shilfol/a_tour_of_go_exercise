package main

import (
	"fmt"
	"math"
)

func Sqrt_ten(x float64) float64 {
	z := 1.0
	for i := 0 ; i < 10 ; i++ {
		z = z - ( z*z - x) / (2 * z)
	}
	return z
}

func Sqrt(x float64) float64 {
	z := 1.0
	tmp := 0.0
	for math.Abs(z - tmp) > 1.0e-13 {
		tmp = z
		z = z - ( z*z - x) / (2 * z)
	}
	return z
}

func main() {
	for i := 1.0 ; i < 10 ; i++ {
		fmt.Println(int(i),Sqrt(i),math.Sqrt(i))
	}
}
