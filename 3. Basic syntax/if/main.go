package main

import (
	"fmt"
)

func CanIDrink(age int) bool {
	//you create variable right before if statement
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println(CanIDrink(15))
}
