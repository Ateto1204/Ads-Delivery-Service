package main

import "fmt"

func swap(a *int, b *int) {
	var tmp int = *a
	*a = *b
	*b = tmp
}

func main() {

	var i, n, x, y int

	fmt.Scanln(&n)

	for i = 0; i < n; i++ {
		fmt.Scanln(&x, &y)
		swap(&x, &y)
		fmt.Println(x, y)
	}

}
