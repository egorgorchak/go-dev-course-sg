package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	slice := []string{"Hi", "Egor", "!"}
	updateSlice(slice)

	fmt.Println(slice)

	a := 4
	ap := &a
	updateInt(ap)

	fmt.Println(a)
	fmt.Println(ap)
}

func updateInt(a *int) {
	*a = 234
}

func updateSlice(slice []string) {
	slice[1] = "Galya"
}

func experimentWithStruct() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			zipCode: 2345,
			email:   "s@s.com",
		},
	}

	jim.updateName("Egor")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	pointerToPerson.firstName = newFirstName
}
