package main

import "fmt"

type contractInfor struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contract  contractInfor
}

func (p person) getFullName() string {
	return p.firstName + " " + p.lastName
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contract: contractInfor{
			email:   "jim@gmail.com",
			zipCode: "94000",
		},
	}
	fmt.Printf("%+v", jim)
}
