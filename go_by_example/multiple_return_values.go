package main

import "fmt"

func vals() (int, int, string) {
	return 3, 7, "test"
}

func main() {

	a, b, _ := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c, _ := vals()
	fmt.Println(c)
}
