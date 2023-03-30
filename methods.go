package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) NoScale(s float64) {
	v.X *= s
	v.Y *= s
}

//uses a pointer
func (v *Vertex) Scale(s float64) {
	v.X *= s
	v.Y *= s
}

func main() {
	v := Vertex{3, 4}

	v.NoScale(10)
	fmt.Println(v)
	
	v.Scale(10)
	fmt.Println(v)
}
