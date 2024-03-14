package main

import "fmt"

type rect struct {
	width  int
	height int
}

// Like making a expanded function for a structure
func (r rect) area() int {
	return r.width * r.height
}

// Declare global variable
var r rect

func main() {

	// Assign structure value to global value
	r = rect{
		width:  5,
		height: 10,
	}

	// Call the structure function
	fmt.Println(r.area())
}
