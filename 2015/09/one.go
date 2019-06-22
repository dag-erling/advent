package main

import (
	"fmt"
)

type City struct {
	name string
	visited bool
}

var cities = map[string]*City{}
var dist = map[*City]map[*City]uint{}

// Find or create a city by name.
func city(name string) *City {
	if cities[name] == nil {
		city := &City{ name: name }
		cities[name] = city
		dist[city] = map[*City]uint{}
	}
	return cities[name]
}

// Find shortest Hamiltonian path.  This is NP-complete, but the data
// set is small enough that a (quadratic) brute force search is
// tractable.
func rfind(a *City) (uint, string) {
	var bestlen uint
	var bestpath string
	for _, b := range cities {
		if ! b.visited {
			b.visited = true
			len, path := rfind(b)
			b.visited = false
			len = dist[a][b] + len
			path = fmt.Sprintf("%d %s %s", dist[a][b], b.name, path)
			if bestlen == 0 || len < bestlen {
				bestlen = len
				bestpath = path
			}
		}
	}
	return bestlen, bestpath
}

func find() (uint, string) {
	var bestlen uint
	var bestpath string
	for _, c := range cities {
		c.visited = true
		len, path := rfind(c)
		c.visited = false
		if bestlen == 0 || len < bestlen {
			bestlen = len
			bestpath = c.name + " " + path
		}
	}
	return bestlen, bestpath
}

func main() {
	for {
		var a, b string
		var d uint
		n, _ := fmt.Scanf("%s to %s = %d\n", &a, &b, &d)
		if n < 3 {
			break
		}
		ca := city(a)
		cb := city(b)
		dist[ca][cb] = d
		dist[cb][ca] = d
	}
	len, path := find()
	fmt.Println(path)
	fmt.Println(len)
}
