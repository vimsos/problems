package dcp

// Write a function, throw_dice(N, faces, total), that determines how many ways it is possible to throw N dice with some number of faces each to get a specific total.
// For example, throw_dice(3, 6, 7) should equal 15.

func throwDice(throws int, faces int, total int) int {
	ways := 0
	var helper func(int, int)

	helper = func(throw int, remaining int) {
		if remaining < 0 {
			return
		}
		if throw == 0 {
			if remaining == 0 {
				ways++
			}
			
			return
		}

		for f := 1; f <= faces; f++ {
			helper(throw-1, remaining-f)
		}
	}

	helper(throws, total)
	return ways
}
