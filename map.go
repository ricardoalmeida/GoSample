package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"foo": 23}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 12
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)
}
