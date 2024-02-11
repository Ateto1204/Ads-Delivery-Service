package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func sol() {
	var x, y int
	fmt.Scan(&x, &y)

	x, y = swap(x, y)
	fmt.Println(x, y)
}

func initTip() {
	fmt.Println("enter 2 integer: ")
}

func main() {
	initTip()
	sol()
}
