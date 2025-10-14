package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if(x > 0) {
		return math.Sqrt(x), nil
	}
	return 0, ErrNegativeTypeSqrt(x)
}

type ErrNegativeTypeSqrt float64;

func (e ErrNegativeTypeSqrt) Error() string {
	return fmt.Sprintf("error trying square root of negative value %v", float64(e))
}

func func20() {
	fmt.Println("func 20: ")

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}