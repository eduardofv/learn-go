package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev_1 := 0
	prev_2 := 0
	
	fn := func() int {
		if prev_1 == 0 && prev_2 == 0 {
			prev_1 = 0
			prev_2 = 1
			return 0
		}
		fib := prev_1 + prev_2
		prev_2 = prev_1
		prev_1 = fib
		return fib
	}
	
	return fn
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

