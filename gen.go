package main

import "fmt"

func InvertDict[K comparable, V comparable](m map[K]V) map[V][]K {
	inv := make(map[V] []K)

	for key, val := range m {
		inv[val] = append(inv[val], key)
	}
	return inv
}

func main() {
	a_map := map[int]string{1: "a", 2: "b", 3: "a"}

	inv := InvertDict(a_map)

	fmt.Println("Org: ", a_map)
	fmt.Println("Inv: ", inv)
}
