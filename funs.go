package main

import "fmt"

func vals() (int, int) {
	return 3, 4
}

func main() {
//	fmt.Println("vals:", vals())
//	fmt.Println("vals: %d %d", vals())
	x,y := vals()
	fmt.Println("vals: %d %d", x, y) //no hace la expansion
//	println("vals:", vals())
	println(vals())
	fmt.Println(vals())
	z := vals()
	println("z:", z)

}
