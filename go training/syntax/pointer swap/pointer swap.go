package main

import "fmt"

func swap(a *int, b *int) {
	tmp := *a // var tmp *int
	*a = *b
	*b = tmp
}

func main() {

	var i, n, x, y int
	const tipkit = "Enter times: " // immutable constant

	fmt.Println(tipkit)

	fmt.Scanln(&n)

	for i = 0; i < n; i++ {
		fmt.Scanln(&x, &y)
		swap(&x, &y)
		fmt.Println(x, y)
	}

}
