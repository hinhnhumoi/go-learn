package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {

	for i := range b {
		b[i] = 'A'
	}

	return 1, nil
}

func func22() {
	fmt.Println("func 22: ")

	reader.Validate(MyReader{})
}