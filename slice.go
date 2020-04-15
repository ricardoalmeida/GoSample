package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}

	fmt.Println(slice)

	slice = append(slice, 4)

	fmt.Println(slice)

	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]

	fmt.Println(s2, s3, s4)
}
