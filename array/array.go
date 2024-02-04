package main

import "fmt"

func main() {
	var i, n int
	var arr [100005]int

	fmt.Println("Enter times: ")
	fmt.Scanln(&n)

	fmt.Println("Enter numbers: ")
	for i = 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Reverse: ")
	for i = n - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
