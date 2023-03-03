package main

import (
	"fmt"
	"time"
)

var gi = 1

func f(from string) {
	gi += 1 
	for i := 0; i < 3; i++ {
		time.Sleep(time.Duration(gi) * time.Second)
		fmt.Println(from, ":", i)
	}
}



func main() {
	go f("goroutine1")
	go f("goroutine2")


	time.Sleep(time.Duration(10) * time.Second)
}
