package main

import "fmt"

func concat(a, b string) string {
	return a + b
}

func sol() {
	var a, b string
	fmt.Scan(&a, &b)

	result := concat(a, b)
	fmt.Println(result)

}

func initTip() {
	fmt.Println("Enter 2 string: ")
}

func main() {
	initTip()
	sol()
}
