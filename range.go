package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}

	for index, num := range nums {
		fmt.Println("index: ", index)
		fmt.Println("num: ", num)
	}

	x := range nums
	fmt.Println(x)
}

