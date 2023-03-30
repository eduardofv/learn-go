package main

import "fmt"

func main() {
	var v interface{}

	v = 10
	fmt.Println(v)

	v = 1.2
	fmt.Println(v)

	v = "hello"
	fmt.Println(v)

}
