package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) person {
	p := person{name: name}
	return p
}

func (self person) have_birthday() {
	self.age += 1
	fmt.Println("in method: ", self)
}

func (self *person) have_birthday_by_ref() {
	self.age += 1
	fmt.Println("in method by ref: ", self)
}

func birthday(who *person) {
	who.age += 1
}

func main() {

	juan := newPerson("juan")

	fmt.Println(juan)

	juan.age = 10

	fmt.Println(juan)

	pjuan := &juan

	fmt.Println(pjuan)
	println(pjuan)
	fmt.Println(*pjuan)

	piter := person{name: "Piter", age: 60}
	fmt.Println(piter)

	birthday(&piter)
	fmt.Println(piter)

	p2 := person(piter)
	fmt.Println("p2:", p2)

	println(&piter)
	println(&p2)

	piter.have_birthday()
	fmt.Println(piter)

	(&piter).have_birthday_by_ref()
	fmt.Println(piter)

}
