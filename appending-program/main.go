package main

import "fmt"

func main() {
	numbers := []int{2, 3}
	numbers = append(numbers, 10)
	fmt.Println(numbers)
}
