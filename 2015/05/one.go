package main

import "fmt"

func is_nice(input string) bool {
	var prev rune
	var have [256]int
	var pair int
	for _, char := range input {
		if char < 'a' || char > 'z' {
			return false
		}
		have[char]++
		if char == prev {
			// pair
			pair++
		}
		if (char == 'b' || char == 'd' || char == 'q' || char == 'y') {
			if char == prev + 1 {
				// forbidden sequence
				return false
			}
		}
		prev = char
	}
	vowels := have['a'] + have['e'] + have['i'] + have['o'] + have['u']
	if vowels < 3 {
		// at least three vowels
		return false
	}
	if pair == 0 {
		// at least one pair
		return false
	}
	return true
}

func main() {
	var input string
	var nice int
	for {
		n, _ := fmt.Scan(&input)
		if n != 1 {
			break
		}
		if is_nice(input) {
			nice++
		}
	}
	fmt.Println(nice)
}
