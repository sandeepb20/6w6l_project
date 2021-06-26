package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var z float64 = 1
	for i := 0; i < 10; i++ {
		fmt.Printf("\nValue of z after %v iteration is %v", i, z)

		z = z - (z*z-x)/(2*z)

	}
	return z
}

func main() {
	z := Sqrt(10)
	fmt.Println("\nFinal Ans is")
	fmt.Println(z)
}
