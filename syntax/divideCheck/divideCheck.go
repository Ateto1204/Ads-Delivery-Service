package main

import (
	"errors"
	"fmt"
	"log"
)

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return dividend / divisor, nil
}

func sol() {
	var a, b float64
	fmt.Scan(&a, &b)

	result, err := divide(a, b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result: ", result)
}

func initTip() {
	fmt.Println("Enter 2 integer: ")
}

func main() {
	initTip()
	sol()
}
