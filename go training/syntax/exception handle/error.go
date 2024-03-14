/**
 *
 * Author: Ateto
 * Description: Exception handling
 *
 */

package main

import "fmt"

type User struct {
	Name string
}

func getUser() (User, error) {
	user := User{Name: "demo"}
	return user, nil
}

func main() {
	user, err := getUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.Name)
}
