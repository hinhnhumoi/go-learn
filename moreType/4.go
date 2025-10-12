package main

import "fmt"

func func4() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9

	fmt.Println(v)
}