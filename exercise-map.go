package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	terms := strings.Fields(s)
	
	for _, term := range terms {
		counts[term] += 1	
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
