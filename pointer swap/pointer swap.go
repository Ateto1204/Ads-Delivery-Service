package main

import "fmt"

func swap(a *int, b *int) {
	var tmp int = *a
	*a = *b
	*b = tmp
}

func main() {
	var x, y int
	fmt.Scanln(&x, &y)
	swap(&x, &y)
	fmt.Println(x, y)
}
