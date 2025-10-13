package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func func4() {
	fmt.Println("Func4: ")
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func func5() {
	fmt.Println("func 5: ")
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}

func func1() {
	fmt.Println("Func1: ")
	v := Vertex{3,4}
	fmt.Println(v.Abs())
}

func func2() {
	fmt.Println("Func2: ")
	v := Vertex{3,4}
	fmt.Println(Abs(v))
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func func6()  {
	fmt.Println("func 6: ")
	v := Vertex{3,4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p:= Vertex{4,3}
	p.Scale(3)
	ScaleFunc(&p, 8)

	fmt.Println(v, p)

}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func func7() {
	fmt.Println("func 7: ")
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

func func8() {
	fmt.Println("func 8: ")
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}



