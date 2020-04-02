package main

import "fmt"

type person struct {
	name     string
	age      int
	fav_food []string
}

func main() {
	//more flexible than map
	fav_food := []string{"pizza", "ramen"}

	//This way of assigning struct variable is undesirable.
	dhk1349 := person{"Daniel", 25, fav_food}

	//This is more clear.
	newdhk1349:=person{name:"Daniel", age:25, fav_food=fav_food}
	fmt.Println(dhk1349)
	fmt.Println(dhk1349.name)
}
