package main

import "fmt"

func main() {

//	a := [2]string{'1', '2'}
	a := [2]int{1,2}

	fmt.Println(a)

	var b [2][3]int

	fmt.Println("len:", len(b))
	fmt.Println("len b[0]:", len(b[0]))
	fmt.Println(b)

	s := make([]string, 3)
	fmt.Println("s:", s)
	s[0] = "hola"
	fmt.Println("s:", s)

	s = append(s, "x")
	fmt.Println("s:", s)

	c := make([]string, len(s))
	fmt.Println("c:", c)
	copy(c, s)
	fmt.Println("c:", c)

	iarr := []string{"a","b","c","d","e"}
	fmt.Println("iarr:", iarr)
	fmt.Println("iarr[2:4]:", iarr[2:4])
	fmt.Println("iarr[0:1]:", iarr[0:1])

	twoD := make([][]int, 3)
	fmt.Println("twoD:", twoD)
	for i:=0; i < 3; i++ {
		twoD[i] = make([]int, i+1)
	}
	fmt.Println("twoD:", twoD)

	subi := iarr[2:4]
	subi[0] = "x"
	fmt.Println("iarr:", iarr)
	fmt.Println("subi:", subi)
	fmt.Println("len(subi):", len(subi))
	fmt.Println("cap(subi):", cap(subi))
	subi[2] = "y"

	//fmt.Println("iarr[-1:]", iarr[-1:])
}
