package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	ret := r.width * r.height
	r.width = 666
	return ret
}

func (r rect) perim() int {
	ret := 2*r.width + 2*r.height
	r.height = 1000
	return ret
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("width outside: ", r.width)
	fmt.Println("perim:", r.perim())
	fmt.Println("height outside: ", r.height)

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
