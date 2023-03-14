package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

func worker(id int) {
	t := rand.Intn(5) + 1
	fmt.Println("Worker", id, "starting, will take", t, "sec")
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Println("Worker", id, "done")
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
