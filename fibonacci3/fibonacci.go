package main

import "fmt"

// fibonacci gives you the sequence of given number
func fibonacci(n int) int {
	if n == 0 || n == 1{
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main(){
	fmt.Println("Fibonacci sequence up to 10:")
	for i :=0; i <= 10; i ++{
		fmt.Println(fibonacci(i))
	}
}