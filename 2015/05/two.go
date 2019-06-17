package main

import "fmt"

func is_nice(input string) bool {
	var p1, p2 rune		// previous two
	var pair int 		// current pair
	var pos [256*256]int	// position of every pair
	var aa_aa, aba bool	// requirements
	for index, char := range input {
		index++	        // convert from zero-based to one-based
		if char < 'a' || char > 'z' {
			return false
		}
		pair = int(p1) * 256 + int(char)
		if pos[pair] == 0 {
			// not seen before, record position
			pos[pair] = index
		} else if index > pos[pair] + 1 {
			// a prior non-overlapping pair exists
			aa_aa = true
		}
		if char == p2 {
			// repeated at a distance of 2
			aba = true
		}
		p2, p1 = p1, char
	}
	return aa_aa && aba
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
