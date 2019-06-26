package main

import (
	"fmt"
)

const SPOONS = 100
const CALORIES = 500

type Ingredient struct {
	name	   string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

var ingredients []*Ingredient

type Cookie struct {
	recipe	   map[*Ingredient]int
	size       int
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
	score      int
}

func (c *Cookie) add(i *Ingredient, n int) {
	c.size += n
	c.recipe[i] += n
	c.capacity += n * i.capacity
	c.durability += n * i.durability
	c.flavor += n * i.flavor
	c.texture += n * i.texture
	c.calories += n * i.calories
	if c.capacity > 0 && c.durability > 0 && c.flavor > 0 && c.texture > 0 {
		c.score = c.capacity * c.durability * c.flavor * c.texture
	} else {
		c.score = 0
	}
}

func (c *Cookie) text() string {
	text := fmt.Sprintf("Score %12d:", c.score)
	for _, i := range ingredients {
		text += fmt.Sprintf(" %3d %s", c.recipe[i], i.name)
	}
	return text
}

var best Cookie

func bake(cookie *Cookie, i int, spoons int) {
	if (i < len(ingredients) - 1) {
		// None of this, but something else
		bake(cookie, i + 1, spoons)
		// Some of this, something else
		for q := 1; q < spoons; q++ {
			cookie.add(ingredients[i], 1)
			bake(cookie, i + 1, spoons - q)
		}
		cookie.add(ingredients[i], 1)
	} else {
		// Some of this, nothing else
		cookie.add(ingredients[i], spoons)
	}
	// fmt.Println(">", cookie.text())
	if cookie.calories == CALORIES && cookie.score > best.score {
		best = *cookie
	}
	cookie.add(ingredients[i], -spoons)
}

func main() {
	ingredients = make([]*Ingredient, 0)
	for {
		i := Ingredient{}
		// Scanf will include the colon in the name
		n, _ := fmt.Scanf("%s capacity %d, durability %d, " +
			"flavor %d, texture %d, calories %d\n",
			&i.name, &i.capacity, &i.durability,
			&i.flavor, &i.texture, &i.calories)
		if n != 6 {
			break
		}
		// Trim the colon
		i.name = i.name[:len(i.name)-1]
		ingredients = append(ingredients, &i)
	}
	bake(&Cookie{ recipe: map[*Ingredient]int{} }, 0, SPOONS)
	fmt.Println(best.score)
}
