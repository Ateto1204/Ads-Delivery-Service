/*
 * Project: 2024 Dcard Backend Intern Assignment
 * Author: Ateto
 *
 * Description: Recommend advertisements to customers
 * 				according to some information of the
 * 				customers.
 *
 * go build *.go
 */

package main

import (
	"net/http"
)

func main() {

	r := setUpRouter()

	http.Handle("/", r)

	// Start the server
	http.ListenAndServe(":8080", nil)
}
