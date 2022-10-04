package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, item := range slice {
		fmt.Printf("%v is %v", item, CalculateEvenOrOdd(item))
		fmt.Println()
	}
}
