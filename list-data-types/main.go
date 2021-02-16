package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	s2 := s1[1:3]
	fmt.Println(s2)
	fmt.Println(s1[2:])
	fmt.Println(s1[:3])
}
