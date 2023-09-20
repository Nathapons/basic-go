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
	var myPerson person
	myPerson.firstName = "Alex"
	myPerson.lastName = "Anderson"
	fmt.Printf("%+v", myPerson)
}
