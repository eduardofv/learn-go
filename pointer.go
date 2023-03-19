package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	v := Point{1, 2}
	fmt.Println(v)

	v.X = 3
	fmt.Println(v)

	pv := &v
	pv.Y = 4
	fmt.Println(v)
}
