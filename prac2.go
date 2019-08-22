package main

import (
	"fmt"
)

func main() {
	//Create a slice of type int and assign 10 values
	x := []int{5, 6, 7, 8, 9, 11, 13, 14, 15, 22}
	//Range over the slice and print the values out
	for k, v := range x {
		fmt.Println(k, v)
	}
	//Using format printing print out the type of the slice
	fmt.Printf("%T\n", x)

}
