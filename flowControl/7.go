package main

import (
	"fmt"
	"math"
)

func pow7(x,n,lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func func7()  {
	fmt.Println(
		pow7(3, 2, 10),
		pow7(3, 3, 20),
	)
}