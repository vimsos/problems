package dcp

// A girl is walking along an apple orchard with a bag in each hand. She likes to pick apples from each tree as she goes along, but is meticulous about not putting different kinds of apples in the same bag.
// Given an input describing the types of apples she will pass on her path, in order, determine the length of the longest portion of her path that consists of just two types of apple trees.
// For example, given the input [2, 1, 2, 3, 3, 1, 3, 5], the longest portion will involve types 1 and 3, with a length of four.

func applePicking(types []int) int {
	start, end, longest, lastIndex := 0, len(types), 0, len(types)

	for end-start > longest {
		length := end - start
		noTypes, typeMap := 0, make(map[int]bool, length)
		for _, t := range types[start:end] {
			if !typeMap[t] {
				noTypes++
				typeMap[t] = true
			}
		}
		if noTypes == 2 && length > longest {
			longest = length
		}
		start++
		end++
		if end > lastIndex {
			end = length - 1
			start = 0
		}
	}

	return longest
}
