package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v\n" , float64(e))
}

func Sqrt(x float64) (float64,error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	} else {
		z := 1.0
		tmp := 0.0
		for math.Abs(z - tmp) > 1.0e-13 {
			tmp = z
			z = z - ( z*z - x) / (2 * z)
		}
		return z , nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
