package main

import "fmt"

func main() {
	var floor int
	var input string
	fmt.Scan(&input)
	for _, char := range input {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	fmt.Println(floor)
}
