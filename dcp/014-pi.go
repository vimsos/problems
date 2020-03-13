package dcp

import (
	"math"
	"math/rand"
	"time"
)

// The area of a circle is defined as πr^2. Estimate π to 3 decimal places using a Monte Carlo method.
// Hint: The basic equation of a circle is x2 + y2 = r2.

func mcpi() float64 {
	rand.Seed(time.Now().UnixNano())

	pi := math.Pi
	var val float64
	hits := 0
	total := 0

	for math.Abs(val-pi) > 0.0001 {
		total++

		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1

		if x*x+y*y <= 1 {
			hits++
		}

		val = 4 * float64(hits) / float64(total)
	}
	return val
}
