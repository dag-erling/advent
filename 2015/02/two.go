package main

import "fmt"

func main() {
	var ribbon int
	for {
		var l, w, h int
		n, _ := fmt.Scanf("%dx%dx%d", &l, &w, &h)
		if (n < 3) { break }
		// half-perimeters of the three sides
		var a, b, c = l + w, w + h, h + l
		// sort them
		if a > b { a, b = b, a }
		if b > c { b, c = c, b }
		if a > b { a, b = b, a }
		// smallest perimeter plus bow (equal to volume)
		ribbon += 2 * a + l * w * h
	}
	fmt.Println(ribbon)
}
