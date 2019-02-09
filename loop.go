package main

import (
	"fmt"
	"math"
)

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	newSum := 1

	for newSum < 1000 {
		newSum += newSum
	}
	fmt.Println(newSum)

	if fin := 1; fin == 0 {
		hang()
	} else {
		fmt.Println(Sqrt(2))
	}

	fmt.Println("FIN")
}

func hang() {
	for {
		fmt.Println("yes")
	}
}

// Sqrt a number
func Sqrt(x float64) float64 {
	z := 1.0
	for math.Abs(z*z-x) > 0.001 {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
