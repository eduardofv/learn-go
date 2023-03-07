package main

import (
	"fmt"
	"time"
)

func tick(name string, done chan bool) {
	ticker := time.NewTicker(500 * time.Millisecond)

	for {
		select {
		case <-done:
			fmt.Println("Stopping ", name)
			ticker.Stop()
			//nomas probando
			//done<-true
			return
		case t := <-ticker.C:
			fmt.Println("Ticker", name, " at: ", t)
		}

	}
}

func main() {

	done := make(chan bool)


	go tick("One", done)
	go tick("Two", done)

	time.Sleep(2000 * time.Millisecond)
	//ticker.Stop()
	done <- true

	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Ticker stopped")
}
