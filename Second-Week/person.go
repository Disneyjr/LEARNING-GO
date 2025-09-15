package main

import "strconv"

type Print interface {
	print() string
}
type Person struct {
	name string
	age  int
}

func Birthday(person *Person) {
	if person != nil {
		println("Happy Birthday")
		person.age = person.age + 1
	}
}
func (person Person) print() string {
	return person.name + " has " + strconv.Itoa(person.age) + " years"
}
func getPrint(print Print) string {
	return print.print()
}
