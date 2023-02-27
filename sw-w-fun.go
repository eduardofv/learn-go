package main

import "fmt"


func make_even(n int) int {
	if n / 2 == 0 {
		return n
	} else {
		return n + 1
	}
}

func main() {
	for i := -3; i <= 3; i++ {
		fmt.Println("Testing ", i)
		switch i {
		case 0: 
			fmt.Println("it's zero")
		case make_even(i):
			fmt.Println("make_even: ", i)
		default: 
			fmt.Println("fall in default")
		}
	}
}
