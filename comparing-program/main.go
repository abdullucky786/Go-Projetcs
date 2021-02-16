package main

import "fmt"

func main() {
	a := map[string]string{"A": "X"}
	b := map[string]string{"A": "y"}
	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)
	if s1 == s2 {
		fmt.Println(" it is equal")
	} else {
		fmt.Println(" it is not equal")
	}
}
