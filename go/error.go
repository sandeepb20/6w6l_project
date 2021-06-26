package main

import (
	"fmt"
)

type ErrNegSqrt float64

func (e ErrNegSqrt) Error() string {

	return fmt.Sprintf("cannot Sqrt negative number: %e", float64(e))
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		err := ErrNegSqrt(x)
		return x, err
	}
	var z float64 = 1
	for i := 0; i < 10; i++ {
		fmt.Printf("\nValue of z after %v iteration is %v", i, z)

		z = z - (z*z-x)/(2*z)

	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
