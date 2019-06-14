package main

import "fmt"

func main() {
	var floor int
	var input string
	fmt.Scan(&input)
	for index, char := range input {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor < 0 {
			fmt.Println(index + 1)
			break
		}
	}
}
