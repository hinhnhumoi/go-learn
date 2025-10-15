package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func func1() {
	fmt.Println("func 1: ")
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}

type List[T any] struct {
	next *List[T]
	val T
}

func (l *List[T]) Push(v T) {
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = &List[T]{val: v}
}

func (l *List[T]) Len() int {
	count := 0
	for n := l; n != nil; n = n.next {
		count++
	}
	return count
}

func (l *List[T]) Values() []T {
	var vals []T
	for n := l; n != nil; n = n.next {
		vals = append(vals, n.val)
	}
	return vals
}

func func2() {
	fmt.Println("func 2: ")
	head := &List[int]{val: 1}
	head.Push(2)
	head.Push(3)
	head.Push(4)

	fmt.Println("List:", head)
	fmt.Println("Length:", head.Len())
	fmt.Println("Values:", head.Values())
}


func main() {
	func1()
	func2()
}
