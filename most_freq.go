package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)


func gen_array(size int, max_int int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(max_int)
	}
	sort.Ints(arr)
	return arr
}

func most_frequent(arr []int) int {
	counts := make(map[int] int)
	var most_freq_count, most_freq_value int;

	for _, value := range arr {
		current_count := counts[value] + 1
		counts[value] = current_count
		if current_count > most_freq_count {
			most_freq_count = current_count
			most_freq_value = value
		}
	}
	return most_freq_value
}

func main() {
	rand.Seed(time.Now().UnixNano())

	arr := gen_array(10, 7)

	fmt.Println("arr: ", arr)
	mfv := most_frequent(arr)
	fmt.Println("most_freq: ", mfv)
}
