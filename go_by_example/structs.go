package main

import "fmt"

type person struct {
	name  string
	age   int
	money float32
}

func NewPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func change(p *person, age int) {
	p.age = age
}

func main() {

	fmt.Println(person{"Bob", 20, 0.0})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(NewPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	s.age = 60
	fmt.Println("modify age 1: ", s.age)
	change(&s, 33)
	fmt.Println("modify age 2: ", s.age)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
