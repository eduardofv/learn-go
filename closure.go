package main

//import "fmt"

func intSeq() func() int {
	i := 0
	f := func() int {
		i++
		return i
	}
	return f
}

func localIntSeq() func() int {
	return func() int {
		x++
		return x
	}
}

func main() {

	f1 := intSeq()

	println(f1())
	println(f1())

	x := 10
	lis := localIntSeq()
	println(lis)
	println(lis)
	println(x)

}
