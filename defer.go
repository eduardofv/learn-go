package main

import "fmt"


func inc(n *int) {
	*n = *n + 1
	fmt.Println("in defer n: ", *n)
}

func afunc() {
	x := 1
	defer inc(&x)

	fmt.Println("x: ", x)
}

func main() {
	afunc()
}
