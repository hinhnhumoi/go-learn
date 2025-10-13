package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func func9() {
	fmt.Println("func 9: ")
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3,4}

	a = f
	a = &v
	a = v
	fmt.Println(a.Abs())
}