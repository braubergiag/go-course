package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 52
	return &p
}
func makePerson(name string) person {
	p := person{name: name}
	p.age = 23
	return p
}
func main() {
	bob := person{"Bob", 20}
	fmt.Println(bob)

	alice := person{name: "Alice", age: 30}
	fmt.Println(alice)

	fred := person{name: "Fred"}
	fmt.Println(fred)

	annptr := &person{name: "Ann", age: 40}
	fmt.Println(annptr)

	john := newPerson("John")
	fmt.Println(john)

	sven := &person{name: "Sven", age: 30}
	fmt.Println((*sven).age)
	fmt.Println(sven.age)
}
