package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	var x, y int
	var visited = map[string]bool{ "0x0": true }
	for _, char := range input {
		switch char {
		case '^':
			y--
		case '<':
			x--
		case 'v':
			y++
		case '>':
			x++
		}
		visited[fmt.Sprintf("%dx%d", x, y)] = true
	}
	fmt.Println(len(visited))
}
