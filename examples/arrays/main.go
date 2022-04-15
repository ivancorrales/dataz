package main

import (
	"fmt"

	"github.com/ivancorrales/dataz"
)

type Person struct {
	firstname string
	age       int
	isFemale  bool
	parent    *Person
}

func (p Person) Gender() string {
	if p.isFemale {
		return "girl"
	}
	return "boy"
}

func (p Person) String() string {
	return fmt.Sprintf("%s is a %v %s years old", p.firstname, p.age, p.Gender())
}

func printDetails(people []Person) {
	dataz.ForEach(people, func(p Person) { fmt.Println(p) })
}

func show(text string) {
	length := len(text)
	fmt.Printf("\n[%s]\n", text)
	for i := 0; i < length; i++ {
		fmt.Print("-")
	}
	fmt.Println("--")
}

func main() {
	studentsClassroomA := []Person{
		{firstname: "John", age: 12, isFemale: false},
		{firstname: "Mary", age: 13, isFemale: true},
		{firstname: "Bert", age: 12, isFemale: false},
		{firstname: "Celia", age: 11, isFemale: true},
	}
	show("Students A")
	printDetails(studentsClassroomA)
	show("Students A (Reverse)")
	output := dataz.Reverse(studentsClassroomA)
	printDetails(output)

	show("Students A (older than 11 years)")
	output = dataz.Filter(studentsClassroomA, func(p Person) bool {
		return p.age > 11
	})
	printDetails(output)

}
