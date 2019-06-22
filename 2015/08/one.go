package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var code, mem int
	lno := 0
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		lno++
		l := len(line)
		code += l
		line = line[1 : l - 1]
		escape := false
		hexcape := 0
		for _, char := range line {
			switch {
			case escape && char == '\\':
				escape = false
				mem++
			case escape && char == '"':
				escape = false
				mem++
			case escape && char == 'x':
				escape = false
				hexcape = 2
				mem++
			case hexcape > 0 && char >= '0' && char <= '9':
				hexcape--
			case hexcape > 0 && char >= 'a' && char <= 'f':
				hexcape--
			case escape || hexcape > 0:
				panic(fmt.Sprintf("invalid sequence on line %d", lno))
			case char == '\\':
				escape = true
			default:
				mem++
			}
		}
	}
	fmt.Println(code - mem)
}
