package main

func main() {
	slice := []int{1, 2, 3}

	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	for _, v := range slice {
		println(v)
	}
}