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

func binsearch_rec(arr []int, value int, offset int) int {
	fmt.Println("----")
	fmt.Println("arr: ", arr)
	fmt.Println("offset: ", offset)

	if len(arr) == 0 {
		return -1
	}

	if len(arr) == 1 {
		if arr[0] == value {
			return offset
		} else {
			return -1
		}
	}
	
	mid := len(arr) / 2
	fmt.Println("mid: ", mid)
	fmt.Println("midval: ", arr[mid])

	if arr[mid] == value {
		return offset + mid
	}

	if value < arr[mid] {
		return binsearch_rec(arr[:mid], value, offset)
	} else {
		return binsearch_rec(arr[mid:], value, offset + mid)
	}
}

func binsearch_it(arr []int, value int) int {
	left := 0
	right := len(arr)
	
	itercount := 0
	fmt.Println("arr: ", arr)
	for left != right {
		fmt.Println("---")
		fmt.Println("left: ", left)
		fmt.Println("right: ", right)
		mid := left + (right - left) / 2
		fmt.Println("mid: ", mid)
		fmt.Println("mid_value: ", arr[mid])
		if arr[mid] == value {
			return mid
		}
		if value < arr[mid] {
			right = mid
		} else { 
			left = mid
		}

		itercount += 1
		if itercount > 20 {
			break
		}
	}
	return -1

}

func test_recursive() {
	arr := gen_array(10, 1000)
	
	//target := -1
	//value := 1010

	target := rand.Intn(len(arr))
	value := arr[target]
	
	fmt.Println(arr)
	fmt.Println("looking for value ", value)
	fmt.Println("looking for index ", target)
	ret := binsearch_rec(arr, value, 0)
	if ret == -1 {
		fmt.Println(">>> not found")
	} else {
		fmt.Println(">>> found index: ", ret)
		fmt.Println(">>> found value: ", arr[ret])
	}
	if target != ret {
		fmt.Print("NO ")
	}
	fmt.Println("OK")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	arr := gen_array(10, 1000)

	target := rand.Intn(len(arr))
	value := arr[target]
	
	fmt.Println(arr)
	fmt.Println("looking for value ", value)
	fmt.Println("looking for index ", target)
	ret := binsearch_it(arr, value)
	if ret == -1 {
		fmt.Println(">>> not found")
	} else {
		fmt.Println(">>> found index: ", ret)
		fmt.Println(">>> found value: ", arr[ret])
	}
	if target != ret {
		fmt.Print("NO ")
	}
	fmt.Println("OK")

//	fmt.Println(gen_array(10, 5))
//	fmt.Println(gen_array(10, 1000))
	
}
