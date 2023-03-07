package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.NewTimer(2 * time.Second)

	<-t1.C //start timer

	for e := range t1 {
		fmt.Println("Timer")
		break
	}
}
