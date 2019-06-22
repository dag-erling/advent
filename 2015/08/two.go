package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var inlen, outlen int
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		inlen += len(line)
		outlen += 2			// surrounding quotes
		for _, char := range line {
			if char == '"' || char == '\\' {
				outlen++	// needs escape
			}
			outlen++
		}
	}
	fmt.Println(outlen - inlen)
}
