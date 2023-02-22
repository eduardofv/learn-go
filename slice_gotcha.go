package main

import ("regexp"
	"io/ioutil" 
	"fmt")

var digits = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	fmt.Println("b:", b)
	fmt.Println("len(b):", len(b))
	fmt.Println("cap(b):", cap(b))
	return digits.Find(b)
}

func main() {
	const fn = "if-fizzbuzz.go"
	d := FindDigits(fn)
	fmt.Println("d:", d)
	fmt.Println("len(d):", len(d))
	fmt.Println("cap(d):", cap(d))
}
