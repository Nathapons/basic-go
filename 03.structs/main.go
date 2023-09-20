package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (p person) getFullName() string {
	return p.firstName + " " + p.lastName
}

func main() {
	var myPerson person = person{firstName: "John", lastName: "Canady"}
	fmt.Println(myPerson.getFullName())
}
