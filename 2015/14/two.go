package main

import (
	"fmt"
)

type Reindeer struct {
	name string
	speed int
	fly int
	rest int
	period int
	distance int
	score int
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
		r.period = r.fly + r.rest
		reindeer[r.name] = &r
	}
	var best *Reindeer
	for t := 0; t < time; t++ {
		furthest := 0
		// fmt.Println("after", t + 1, "seconds")
		for _, r := range reindeer {
			if t % r.period < r.fly {
				// currently flying
				r.distance += r.speed
			}
			if r.distance > furthest {
				furthest = r.distance
			}
			// fmt.Println(r.name, "has flown", r.distance, "km")
		}
		for _, r := range reindeer {
			if r.distance == furthest {
				r.score++
				if best == nil || r.score > best.score {
					best = r
				}
			}
			// fmt.Println(r.name, "has", r.score, "points")
		}
	}
	fmt.Println(best.score)
}
