package main

import "fmt"

func main() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	fmt.Println(u)

	u.ID = 1
	u.FirstName = "Ricardo"
	u.LastName = "Almeida"
	fmt.Println(u)

	u2 := user{ID: 1, FirstName: "Jao", LastName: "Boeira"}
	fmt.Println(u2)
}
