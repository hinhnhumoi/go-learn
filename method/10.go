package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func func10() {
	fmt.Println("func 10: ")

	var i I = T{"Hello"}
	i.M()
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func func11() {
	fmt.Println("func 11: ")
	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

type I2 interface {
	M2()
}

func describe2(i I2) {
	fmt.Printf("(%v, %T)\n", i, i)
}


func (t *T) M2() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func func12() {
	fmt.Println("func 12: ")

	var i I2

	var t *T
	i = t
	describe2(i)
	i.M2()

	i = &T{"hello"}
	describe2(i)
	i.M2()
}

func func14() {
	fmt.Println("func14: ")
	var i interface{}
	describe3(i)

	i = 42
	describe3(i)

	i = "hello"
	describe3(i)

}

func describe3(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func func15() {
	fmt.Println("func 15: ")

	var i interface{} = "hello"

	describe3(i)

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

}
