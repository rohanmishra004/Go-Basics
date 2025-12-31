package main

import "fmt"

//We setup structs to work as receiver functions too

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Halper",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jim.updateName("Jimmy")
	jim.print()

}

//the person struct is a receiver here
func (p person) print() {
	fmt.Printf("%+v", p)

}

//we are going to write a function called update name which updates the person as the receiver

//This will not directly update the name of the object , so here we will be introduced to pointers
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
