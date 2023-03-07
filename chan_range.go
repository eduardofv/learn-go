package main

import "fmt"

func main() {

	q := make(chan string, 2)

	q <- "one"
	q <- "two"
//	q <- "three"

	close(q)

	for e := range q {
		fmt.Println(e)
	}
}
