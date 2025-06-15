package main

import "math"

func asteroidCollision(asteroids []int) []int {
    
	if len(asteroids) == 0 {

		return []int{}
	}

	Collisions := []int{}

	for _,asteroid := range asteroids {

		if asteroid > 0 {

			Collisions = append(Collisions, asteroid)
		}else {

			for len(Collisions) > 0 && Collisions[len(Collisions) - 1] > 0 && Collisions[len(Collisions) - 1] < int(math.Abs(float64(asteroid))) {

				Collisions = Collisions[:len(Collisions) - 1]
			}

			if len(Collisions) == 0 || Collisions[len(Collisions) - 1] < 0 {

				Collisions = append(Collisions, asteroid)
			}else if Collisions[len(Collisions) - 1] == int(math.Abs(float64(asteroid))) {

				Collisions = Collisions[:len(Collisions) - 1]
			}
		}
	}

	return Collisions
}