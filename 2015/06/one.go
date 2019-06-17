package main

import "fmt"

type LightGrid struct {
	light [1000][1000]bool
}

func turn(grid *LightGrid, x1, y1, x2, y2 int, onoff bool) {
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			grid.light[y][x] = onoff
		}
	}
}

func toggle(grid *LightGrid, x1, y1, x2, y2 int) {
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			grid.light[y][x] = ! grid.light[y][x]
		}
	}
}

func main() {
	var grid LightGrid
	var word string
	var n, x1, y1, x2, y2 int
	for {
		n, _ = fmt.Scan(&word)
		if n != 1 {
			// end of input
			break
		}
		switch word {
		case "turn":
			n, _ = fmt.Scanf("%s %d,%d through %d,%d\n",
				&word, &x1, &y1, &x2, &y2)
			if n == 5 {
				switch word {
				case "on":
					turn(&grid, x1, y1, x2, y2, true)
				case "off":
					turn(&grid, x1, y1, x2, y2, false)
				}
			}
		case "toggle":
			n, _ = fmt.Scanf("%d,%d through %d,%d\n",
				&x1, &y1, &x2, &y2)
			if n == 4 {
				toggle(&grid, x1, y1, x2, y2)
			}
		default:
			// invalid input
			break
		}
	}
	// count outcome
	var count int
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			if grid.light[y][x] {
				count++
			}
		}
	}
	fmt.Println(count)
}
