package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := 0; y < dy ; y++ {
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			row[x] = uint8((x+y)/2)
		}
		pic[y] = row
	}
	return pic
}

func func18()  {
	fmt.Println("func 18: ")
	pic.Show(Pic)
}