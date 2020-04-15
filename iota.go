package main

import "fmt"

const (
	first  = iota + 6
	second = 2 << iota
	a      = 2 << iota
	b      = 2 << iota
	c      = 2 << iota
)

func main() {
	fmt.Println(first, second)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	arr := [3]int{1, 2, 3}

	fmt.Println(arr[1])
}
