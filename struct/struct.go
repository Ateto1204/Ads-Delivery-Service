package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	var n, i int

	fmt.Println("Enter times: ")
	fmt.Scanln(&n)

	for i = 0; i < n; i++ {
		person := Person{}
		fmt.Scan(&person.name, &person.age)
		fmt.Printf("My name is %s and I'm %d\n", person.name, person.age)
	}
}
