package main

import "fmt"

const x = 10

func main() {
	for i := 0; i < 10; i += 2 {
		fmt.Println(i * x)
	}
}
