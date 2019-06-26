package main

import (
	"fmt"
)

type Reindeer struct {
	name string
	speed int
	fly int
	rest int
}

func main() {
	time := 0
	fmt.Scanf("%d", &time)
	reindeer := map[string]*Reindeer{}
	for {
		r := Reindeer{}
		n, _ := fmt.Scanf("%s can fly %d km/s for %d seconds, " +
			"but then must rest for %d seconds.\n",
			&r.name, &r.speed, &r.fly, &r.rest)
		if n != 4 {
			break
		}
		reindeer[r.name] = &r
	}
	best := 0
	for _, r := range reindeer {
		// fmt.Printf("%s can fly %d km/s for %d seconds, " +
		// 	"but then must rest for %d seconds.\n",
		// 	r.name, r.speed, r.fly, r.rest)
		period := r.fly + r.rest
		q := time / period	// complete periods
		m := time % period	// final incomplete period
		if m > r.fly {
			m = r.fly
		}
		dist := (q * r.fly + m) * r.speed
		// fmt.Printf("%s flies %d full %d-second periods plus " +
		// 	"%d seconds for a total of %d km.\n",
		// 	r.name, q, period, m, dist)
		if dist > best {
			best = dist
		}
	}
	fmt.Println(best)
}
