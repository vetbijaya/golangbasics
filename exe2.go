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

	//store the values of type person in a map with the key of last name

	m := map[string]person{
		p1.lastname: p1,
		p2.lastname: p2,
	}

	//Print out the values, ranging over the slice

	for _, v := range m {
		fmt.Println(v.firstname)
		fmt.Println(v.lastname)
		for k, v1 := range v.favIcecreamFlavors {
			fmt.Println(k, v1)
		}
		fmt.Println("------------")
	}

}
