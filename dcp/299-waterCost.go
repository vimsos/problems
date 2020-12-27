package dcp

// A group of houses is connected to the main water plant by means of a set of pipes. A house can either be connected by a set of pipes extending directly to the plant, or indirectly by a pipe to a nearby house which is otherwise connected.
// For example, here is a possible configuration, where A, B, and C are houses, and arrows represent pipes:
// A <--> B <--> C <--> plant
// Each pipe has an associated cost, which the utility company would like to minimize. Given an undirected graph of pipe connections, return the lowest cost configuration of pipes such that each house has access to water.
// In the following setup, for example, we can remove all but the pipes from plant to A, plant to B, and B to C, for a total cost of 16.
// pipes = {
//     'plant': {'A': 1, 'B': 5, 'C': 20},
//     'A': {'C': 15},
//     'B': {'C': 10},
//     'C': {}
// }

func waterCost(pipeSet map[string]map[string]uint) map[string]map[string]uint {
	type pipe struct {
		a, b string
		cost uint
	}
	calcCost := func(set map[string]map[string]uint) uint {
		cost := uint(0)
		for _, place := range set {
			for _, c := range place {
				cost += c
			}
		}
		return cost
	}
	rebuild := func(pipes []pipe) map[string]map[string]uint {
		set := map[string]map[string]uint{}
		for _, p := range pipes {
			if set[p.a] == nil {
				set[p.a] = map[string]uint{}
			}
			set[p.a][p.b] = p.cost
		}
		return set
	}
	checkBit := func(n int, pos int) bool {
		val := n & (1 << pos)
		return (val > 0)
	}

	// derive a list of edges and a list of unique house names from the graph
	pipes, houses := []pipe{}, []string{}
	for a, links := range pipeSet {
		if a != "plant" {
			houses = append(houses, a)
		}
		for b, c := range links {
			pipes = append(pipes, pipe{a: a, b: b, cost: c})
		}
	}

	// define the baseline cost for comparison
	minCost, minPipes, setLength := calcCost(pipeSet), pipeSet, len(pipes)

	// iterate through the subsets of all edges
	totalSets := 1 << setLength
	for set := 0; set < totalSets; set++ {
		candidateSet := []pipe{}
		for position := 0; position < setLength; position++ {
			if checkBit(set, position) {
				candidateSet = append(candidateSet, pipes[position])
			}
		}

		// rebuild the graph from the reduced set of edges
		candidate := rebuild(candidateSet)
		candidateCost := calcCost(candidate)

		// bail early if the candidate cost is greater than the current minimum cost
		if candidateCost > minCost {
			continue
		}

		// do a breadth-first traversal of the graph, verify that all houses are reachable starting from the water plant
		toVisit := make(chan string, setLength)
		toVisit <- "plant"
		visited := map[string]bool{}
		for next := range toVisit {
			visited[next] = true
			for house := range candidate[next] {
				if !visited[house] {
					toVisit <- house
				}
			}
			if len(toVisit) == 0 {
				close(toVisit)
			}
		}

		// check if all houses are reachable in this subset
		allHousesAreReachable := true
		for _, h := range houses {
			allHousesAreReachable = visited[h] && allHousesAreReachable
		}
		if allHousesAreReachable {
			minPipes = candidate
			minCost = candidateCost
		}
	}

	return minPipes
}
