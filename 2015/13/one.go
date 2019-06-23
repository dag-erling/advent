package main

import (
	"fmt"
	"strings"
)

var hap = map[string]map[string]int{}
var v = map[string]bool{}
var t []string
var n int
var best int

// This is similar to a Hamiltonian cycle, which is NP-complete, but
// the data set is small enough for a brute-force search.

func find(i, sum int) {
	if i == n {
		sum += hap[t[n-1]][t[0]] + hap[t[0]][t[n-1]]
		if sum > best {
			best = sum
//			fmt.Printf("New best sum: %d", sum)
//			for _, s := range t {
//				fmt.Printf(" %s", s)
//			}
//			fmt.Printf("\n")
		}
		return
	}
	for a, _ := range hap {
		if ! v[a] {
			v[a] = true
			t[i] = a
			find(i + 1, sum + hap[t[i-1]][a] + hap[a][t[i-1]])
			v[a] = false
		}
	}
}

func main() {
	var lno int
read:
	for {
		var a, b, gl string
		var score int
		// Note that Scanf() only stops at whitespace, we will
		// have to trim the period from the second name later.
		n, _ := fmt.Scanf("%s would %s %d happiness units by sitting next to %s\n",
			&a, &gl, &score, &b)
		lno++
		switch n {
		case 0:
			break read
		case 4:
			// ok
		default:
			panic(fmt.Sprintf("Invalid input on line %d", lno))
		}
		switch gl {
		case "gain":
		case "lose":
			score = -score
		default:
			panic(fmt.Sprintf("Invalid input on line %d", lno))
		}
		b = strings.TrimSuffix(b, ".")
		if hap[a] == nil {
			hap[a] = map[string]int{}
		}
		hap[a][b] = score
	}
	n = len(hap)
	t = make([]string, n)
	for a, _ := range hap {
		v[a] = true
		t[0] = a
		find(1, 0)
		v[a] = false
	}
	fmt.Println(best)
}
