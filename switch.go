package main

import (
	"fmt"
//	"time"
)

func main() {
	s := "hola"

	switch s {
		case "adios": 
			fmt.Println("Adios")
		case "hola":
			fmt.Println("Hola")
		default:
			fmt.Println("What?")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("arg is bool")
		default:
			fmt.Println("arg is not bool, it is %T\n", t)
		}
	}

	whatAmI(true)
}

