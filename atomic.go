package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		i := i
		go func() {
			for c := 0; c < 3; c++ {
				fmt.Println("worker ", i)
				atomic.AddUint64(&ops, 1)
				fmt.Println(i, "ops: ", ops)
				time.Sleep(3 * time.Second)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops)
}
