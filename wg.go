package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

func worker(id int) {
	fmt.Println("Worker ", id, " starting")
	time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	fmt.Println("Worker ", id, "done")
}

func main() {

	var wg sync.WaitGroup

	for i:=1; i <= 10; i++ {
		wg.Add(1)

		i := i //ver FAQ, esta del nabo

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	fmt.Println("Waiting")
	wg.Wait()
	fmt.Println("Finished")
}
