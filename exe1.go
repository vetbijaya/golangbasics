// struct exercise on August 22, 2019
package main

import (
	"fmt"
)

//declare the struct
type person struct {
	firstname          string
	lastname           string
	favIcecreamFlavors []string
}

func main() {
	// create two values of type person
	p1 := person{
		firstname:          "David",
		lastname:           "Collin",
		favIcecreamFlavors: []string{"vanilla", "strawberry"},
	}
	p2 := person{
		firstname:          "Miley",
		lastname:           "Perry",
		favIcecreamFlavors: []string{"chocolate", "butter pecan"},
	}
	// print out the values, ranging over the elements in slice that stores fav colors

	fmt.Println(p1.firstname)
	fmt.Println(p1.lastname)
	for k, v := range p1.favIcecreamFlavors {
		fmt.Println(k, v)
	}
	fmt.Println(p2.firstname)
	fmt.Println(p2.lastname)
	for k, v := range p2.favIcecreamFlavors {
		fmt.Println(k, v)
	}
}
