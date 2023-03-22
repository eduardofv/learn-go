package main

import "fmt"

func build_closures() (func(int) int, func(int) int) {
	val := 0

	add := func(x int) int {
		val += x
		return val
	}

	diff := func(x int) int {
		val -= x
		return val
	}

	return add, diff
}

func main() {
	adder, differ := build_closures()

	fmt.Println("Adding 10: ", adder(10))
	fmt.Println("Adding 10: ", adder(10))
	fmt.Println("Adding 10: ", adder(10))
	fmt.Println("Diff 1: ", differ(1))
	fmt.Println("Adding 10: ", adder(10))
	fmt.Println("diff 1: ", differ(1))
}
