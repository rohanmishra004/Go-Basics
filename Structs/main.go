package main

import "fmt"

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
	jimPointer := &jim             //this allows us to point to the memory address for the jim
	jimPointer.updateName("Jimmy") //since jimPointer now points to the memory now we are able to update the value so *jimPointer is able to update the value of the struct
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName

}
