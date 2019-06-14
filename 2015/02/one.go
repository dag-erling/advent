package main

import "fmt"

func main() {
	var paper int
	for {
		var l, w, h int
		n, _ := fmt.Scanf("%dx%dx%d", &l, &w, &h)
		if (n < 3) { break }
		fmt.Println(l, w, h)
		// areas of the three sides
		var a, b, c = l * w, w * h, h * l
		// sort them
		if a > b { a, b = b, a }
		if b > c { b, c = c, b }
		if a > b { a, b = b, a }
		// two of each plus slop
		paper += 2 * a + 2 * b + 2 * c + a
	}
	fmt.Println(paper)
}
