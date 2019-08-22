package main

import (
	"fmt"
)

func main() {
	//Create an array which holds 5 values of type int
	x := [5]int{5, 6, 7, 8, 9}

	//Range over the array and print out the values using format prining
	for k, v := range x {
		fmt.Println(k, v)
	}
	// print out the type of the array
	fmt.Printf("%T\n", x)
}
