package main

import "fmt"

func getName(firstName, secondName string) (string, string) {
	return firstName, secondName
}

func sol() {
	var firstName, secondName string
	fmt.Scan(&firstName, &secondName)

	resultName, _ := getName(firstName, secondName)
	fmt.Println("Result: " + resultName)
}

func initTip() {
	fmt.Println("Enter 2 string: ")
}

func main() {
	initTip()
	sol()
}
