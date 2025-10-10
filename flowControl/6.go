package main

import (
	"fmt"
	"math"
)

func pow6( x, n, lim float64) float64 {
	if v:= math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func func6() {
	fmt.Println(pow6(3, 2, 10), pow6(3, 3, 20))
}