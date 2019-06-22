package main

import (
	"fmt"
)

// Note: repeated conversion between strings and byte arrays is
// extremely inefficient.

// Increment a string
func increment(s string) string {
	b := []byte(s)
	var i int
	for i = len(b) - 1; i >= 0; i-- {
		if b[i] < 'z' {
			b[i]++
			// shortcut: skip illegal characters
			if b[i] == 'i' || b[i] == 'l' || b[i] == 'o' {
				b[i]++
			}
			break
		} else {
			b[i] = 'a'
		}
	}
	s = string(b)
	if i < 0 {
		s = "a" + s
	}
	return s
}

func check(s string) bool {
	var d1, d2, t byte
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == s[i - 1] {
			if d1 == 0 {
				d1 = s[i]
				// fmt.Printf("first pair: %c%c\n", d1, d1)
			} else if d2 == 0 && s[i] != d1 {
				d2 = s[i]
				// fmt.Printf("second pair: %c%c\n", d2, d2)
			}
		}
		if i > 2 && s[i] == s[i-1] + 1 && s[i] == s[i-2] + 2 {
			t = s[i]
			// fmt.Printf("triplet: %c%c%c\n", s[i-2], s[i-1], s[i])
		}
	}
	return d1 != 0 && d2 != 0 && t != 0
}

func next(pwd string) string {
	for {
		pwd = increment(pwd)
		if check(pwd) {
			return pwd
		}
	}
}

func main() {
	var pwd string
	fmt.Scan(&pwd)
	fmt.Println(next(next(pwd)))
}
