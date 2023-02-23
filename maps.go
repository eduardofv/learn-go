package main

import "fmt"

func main() {
	var intmap = make(map[int]int)

	intmap[0] = 1
	x := intmap[-1]
	intmap[-2] += 1


	println(intmap[0])
	println(x)
	println(intmap[-2])


	k,v := intmap[0]
	println("k:", k)
	println("v:", v)

	k,v = intmap[-3]
	println("k:", k)
	println("v:", v)

	strmap := map[string]int{"foo": 1, "bar": 2}
	println(strmap)
	fmt.Println(strmap)

}

