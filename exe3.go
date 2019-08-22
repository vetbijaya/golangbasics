//embedded struct

package main

import (
	"fmt"
)

// Create a new type vehicle
type vehicle struct {
	doors int
	color string
}

//Create two new types: truck & sedan
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		luxury: false,
	}

	//Print out eack of these values
	fmt.Println(t)
	fmt.Println(s)
	//Print out a single field from each of these values
	fmt.Println(t.color)
	fmt.Println(s.color)
}
