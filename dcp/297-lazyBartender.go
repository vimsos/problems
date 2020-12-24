package dcp

// At a popular bar, each customer has a set of favorite drinks, and will happily accept any drink among this set. For example, in the following situation, customer 0 will be satisfied with drinks 0, 1, 3, or 6.
// preferences = {
//     0: [0, 1, 3, 6],
//     1: [1, 4, 7],
//     2: [2, 4, 7, 5],
//     3: [3, 2, 5],
//     4: [5, 8]
// }
// A lazy bartender working at this bar is trying to reduce his effort by limiting the drink recipes he must memorize. Given a dictionary input such as the one above, return the fewest number of drinks he must learn in order to satisfy all customers.
// For the input above, the answer would be 2, as drinks 1 and 5 will satisfy everyone.

func lazyBartender(preferences map[int][]int) int {
	contains := func(slice []int, item int) bool {
		for _, current := range slice {
			if item == current {
				return true
			}
		}
		return false
	}

	// create a list of unique drink indexes
	drinkList := []int{}
	{
		drinkMap := make(map[int]bool)
		for _, favorites := range preferences {
			for _, drink := range favorites {
				if _, exists := drinkMap[drink]; !exists {
					drinkList = append(drinkList, drink)
				}
				drinkMap[drink] = true
			}
		}
	}

	// create a power set of drinks indexes
	drinkSet := powerSet(drinkList)
	minimum := len(drinkList)

	// exhaustively check for a subset that satisfies every customers
	for _, set := range drinkSet {
		// bail early if set is empty or bigger than the current minimum
		if len(set) >= minimum || len(set) == 0 {
			continue
		}

		customersHaveDrink := make(map[int]bool)
		for _, drink := range set {
			for customer, favorites := range preferences {
				customersHaveDrink[customer] = contains(favorites, drink) || customersHaveDrink[customer]
			}
		}

		// a set is valid if all customers have at least one favorite drink inside the set
		setIsValid := true
		for _, customerHasDrink := range customersHaveDrink {
			setIsValid = setIsValid && customerHasDrink
		}

		if setIsValid {
			minimum = len(set)
		}
	}

	return minimum
}
