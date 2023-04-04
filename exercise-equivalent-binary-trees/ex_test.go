package main

import (
	"golang.org/x/tour/tree"
	"testing"
)


//solo para probar como correr una prueba
func TestWalk(t *testing.T) {
	tree := tree.New(1)
	ch := make(chan int, 1)

	go Walk(tree, ch)
	ret := <- ch

	if ret != 1 { //forzar que falle con 2
		t.Fatalf("Failed with value %v", ret)
	}
}
