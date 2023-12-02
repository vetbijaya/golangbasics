package main

import "fmt"

// fibonacci gets the nth fibonacci number
func fibonacci(n int) int {
	if n < 2{
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main(){
	fmt.Println("The fibonacci of 7 is:", fibonacci(7)) //output is "The fibonacci of 7 is: 13"
}