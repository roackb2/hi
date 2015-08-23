package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	skills []string
}

func NewPerson() (person Person) {
	person = Person{age: 0, name: "Hello World", skills: []string{"html", "css"}}
	return
}

func (person *Person) Print() {
	fmt.Println("name: ", person.name)
	fmt.Println("age: ", person.age)
}

func (p Person) ListSkills() {
	for i := range p.skills {
		fmt.Printf("%s could write %s\n", p.name, p.skills[i])
	}
}
