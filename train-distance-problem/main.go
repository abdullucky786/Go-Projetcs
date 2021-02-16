package main

import "fmt"

func main() {
	fmt.Println("enter the distance:")
	var distance int
	fmt.Scanln(&distance)
	fmt.Println("enter the speed:")
	var speed int
	fmt.Scanln(&speed)
	var time = distance / speed
	fmt.Println("the time of train:", time)
}
