package main

import "fmt"

func main() {
	names := [5]string{"nico", "lynn", "danny"}
	fmt.Println(names)
	names[3] = "three"
	names[4] = "four"
	fmt.Println(names)

	//slice object is assigned like below.
	//slice doesn't have limitation in space(?).
	slicevar := []string{"1", "2", "3"}
	//function append doesn't revise object
	//but returns new result using fnc append.
	slicevar = append(slicevar, "4")
	fmt.Println(slicevar)
}
