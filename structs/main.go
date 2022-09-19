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

// give me the value this memory address is pointing at
func (ptp *person) updateName(f string) { // this * is saying that it only accepts the pointer to the object
	ptp.firstName = f
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	// jin := person{firstName: "Jin", lastName: "Kim"}

	// fmt.Printf("first name is %s and last name is %s", jin.firstName, jin.lastName)

	// 	var jin person

	// 	jin.firstName = "Jin"
	// 	jin.lastName = "Kim"

	jin := person{
		firstName: "Jin",
		lastName:  "Kim",
		contact: contactInfo{
			email:   "abc@gmail.com",
			zipCode: 00001,
		},
	}

	//& is an operator it gives memory address of the value this variable is pointing at
	// jimPointer := &jin
	jin.updateName("Jinfluenza")
	jin.print()
}
