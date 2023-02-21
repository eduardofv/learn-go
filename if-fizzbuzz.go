package main

import "fmt"

func main() {
	for i:=0; i<100; i++ {
		d3 := i % 3
		d5 := i % 5 
		if d3 == 0 { fmt.Print("fizz") }
		if d5 == 0 { fmt.Print("buzz") }
		if d3 != 0 && d5 != 0 { fmt.Print(i) }
		fmt.Println()
	}
}


