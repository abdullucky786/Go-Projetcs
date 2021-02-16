package main

import "fmt"

func main() {
	fmt.Println("Enter the First number:")
	var num1 int
	fmt.Scanln(&num1)
	fmt.Println("Enter the Second number:")
	var num2 int
	fmt.Scanln(&num2)
	var sum = num1 + num2
	fmt.Printf("sum of the two numbers:", sum)
}
