package main

import "fmt"

func main() {
	a := 2
	b := &a
	a = 5
	*b = 500 //value can be chagned using *.
	fmt.Println(a, b)
	// * looks through the addr
	fmt.Println(*b)
	fmt.Println(a)

	//* works just as same as c++'s *.
}
