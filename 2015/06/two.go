package main

import "fmt"

type LightGrid struct {
	light [1000][1000]int
}

func turn_on(grid *LightGrid, x1, y1, x2, y2 int) {
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			grid.light[y][x] += 1
		}
	}
}

func turn_off(grid *LightGrid, x1, y1, x2, y2 int) {
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			if grid.light[y][x] > 0 {
				grid.light[y][x] -= 1
			}
		}
	}
}

func toggle(grid *LightGrid, x1, y1, x2, y2 int) {
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			grid.light[y][x] += 2
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
					turn_on(&grid, x1, y1, x2, y2)
				case "off":
					turn_off(&grid, x1, y1, x2, y2)
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
	var brightness int
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			brightness += grid.light[y][x]
		}
	}
	fmt.Println(brightness)
}
