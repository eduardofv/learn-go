//https://go.dev/tour/concurrency/8
package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}


// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make (chan int, 10)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	for i := 0; i < 10; i++ {
		x := <- ch1
		y := <- ch2
		if x != y {
			return false
		}
		/*equivs:
		if x, y := <- ch1, <- ch2; x!=y {
			return false
		}
		if  <- ch1 != <- ch2 {
			return false
		}
		*/
	}
	return true
}

func main() {
	ch := make(chan int, 10)
	
	t := tree.New(1)
	fmt.Println("Tree: ", t.String())
	go Walk(tree.New(1), ch)
	defer close(ch)
	
	for i := 0; i < 10; i++ {
		x := <- ch
		fmt.Println(x)
	}
	
	fmt.Println(Same(tree.New(1), tree.New(1)))
}

