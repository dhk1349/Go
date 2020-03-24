package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) { //if you want to get multiple parameters that has same type, use ... prior to type name.
	fmt.Println(words)
}

func naked_lenAndUpper(name string) (length int, uppercase string) {
	//defer is line that is executed after function finishes(returns sth).
	//This can be done in such manners: deleting picture after using it, sending API to server...
	defer fmt.Println("I'm done")
	//This is naked form of return.
	//In naked form, variables to return are expressed in the upper line.
	length = len(name) //since var length is initialized in line 20, we don't use := to assign value.
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	/*
		for i:=0;i<len(numbers); i++{
			you can code for loop like this too
		}
	*/
	return total
}

func main() {
	//var name string = "dhk1349"
	name := "dhk1349" //Go will guess the type of variable automatically. But cannot change type of the variable later on.
	name = "new_dhk1349"
	fmt.Println(name)
	totallen, uppername := naked_lenAndUpper(name)
	//totallen, _ := lenAndUpper(name) you can use underscore letter if you are not using returned value.
	fmt.Println(totallen, uppername)
	//fmt.Println(multiply(2, 2))

	repeatMe("nico", "daniel", "dhk1349")

	fmt.Println(superAdd(10, 1, 2, 3, 4))
}
