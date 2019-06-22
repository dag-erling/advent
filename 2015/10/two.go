package main

import (
	"fmt"
)

func look_say(input string) (output string) {
	var prev rune
	var count uint
	for _, char := range input + "@" {
		switch {
		case prev == 0:
			prev = char
			count = 1
		case prev == char:
			count++
		default:
			output += fmt.Sprintf("%d%c", count, prev)
			prev = char
			count = 1
		}
	}
	return
}

func main() {
	var sequence string
	fmt.Scan(&sequence)
	for i := 1; i <= 50; i++ {
		// fmt.Printf("%d ", i)
		sequence = look_say(sequence)
		// fmt.Printf("%d\n", len(sequence))
	}
	fmt.Println(len(sequence))
}
