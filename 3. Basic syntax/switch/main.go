package main

import (
	"fmt"
)

func CanIDrink(age int) bool {
	//you create variable right before if statement
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 90:
		return false
	}
	/*
		you can code switch like this too.
		switch koreanage:=age+2; koreanage{
		case 18:
			return false
		case 17:
			return true
		case 90:
			return false
		}
	*/
	return false
}

func main() {
	fmt.Println(CanIDrink(15))
}
