package main

import "fmt"

func main() {
	var i, j int
	i, j = 5, 8
	j, i = i, j
	fmt.Println(i, j)
}
