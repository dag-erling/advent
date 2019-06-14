package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	var x, y, rx, ry int
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
		// Swap between Santa and Robo-Santa
		x, y, rx, ry = rx, ry, x, y
	}
	fmt.Println(len(visited))
}
