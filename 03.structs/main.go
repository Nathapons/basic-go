package main

import "fmt"

type contractInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contractInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contractInfo: contractInfo{
			email:   "jim@gmail.com",
			zipCode: "94000",
		},
	}
	jim.updateName("Jimmy")
	jim.print()
}
