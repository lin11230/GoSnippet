package main

import "fmt"

type Rectangle struct {
	height, width float64
}

func area(r Rectangle) float64 {
	return r.height * r.width
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{3, 9}

	fmt.Println("Area of r1 is: ", area(r1))
	fmt.Println("Area of r2 is: ", area(r2))

}
