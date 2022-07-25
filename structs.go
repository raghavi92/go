package main

import "fmt"

type User struct {
	firstName string
	lastName  string
	state     string
}

func main() {
	user := User{
		firstName: "Alice",
		lastName:  "Jacob",
		state:     "Active",
	}

	user.print()
	user.updateState("Inactive")
	user.print()
}

func (u User) print() {
	fmt.Printf("%+v", u)
}

func (u *User) updateState(newState string) {
	u.state = newState
}
