package main

import (
	"fmt"
	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Piter")
	fmt.Println(message)
}

