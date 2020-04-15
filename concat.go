package main

import (
	"fmt"
)

func main() {
	var firstName *string = new(string)

	*firstName = "Ricardo"
	fmt.Println(fmt.Sprintf("Hello Gophers! %s", *firstName))
}
