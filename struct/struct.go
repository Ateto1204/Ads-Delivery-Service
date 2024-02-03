package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	var n, i, person_age int
	var person_name string

	fmt.Println("Enter times: ")
	fmt.Scanln(&n)

	for i = 0; i < n; i++ {
		fmt.Scanln(&person_name, &person_age)
		var person Person = Person{name: person_name, age: person_age}
		fmt.Printf("My name is %s and I'm %d\n", person.name, person.age)
	}
}
