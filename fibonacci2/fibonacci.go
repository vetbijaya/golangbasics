package main

import "fmt"

// generateFibonacci function generates the first n numbers in the Fibonacci sequence and returns them as a slice
func generateFibonacci(n int) []int {
	fibonacci := make([]int, n)
	for i :=0; i < n; i++{
		if i <= 1 {
			fibonacci[i] = i
		} else {
			fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
		}
	}
	return fibonacci
}

func main(){
	result := generateFibonacci(5)
	fmt.Println(result) // Output: [0 1 1 2 3]
}