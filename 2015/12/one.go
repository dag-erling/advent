package main

import (
	"fmt"
)

func main() {
	var json string
	fmt.Scan(&json)
	var num, sum int
	var neg bool
	for _, char := range json {
		switch {
		case char == '-':
			neg = true
		case char >= '0' && char <= '9':
			num = num * 10 + int(char) - '0'
		default:
			if neg {
				sum -= num
			} else {
				sum += num
			}
			neg = false
			num = 0
		}
	}
	fmt.Println(sum)
}
