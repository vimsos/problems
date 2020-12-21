package dcp

// Pascal's triangle is a triangular array of integers constructed with the following formula:
// The first row consists of the number 1.
// For each subsequent row, each element is the sum of the numbers directly above it, on either side.
// For example, here are the first few rows:
//     1
//    1 1
//   1 2 1
//  1 3 3 1
// 1 4 6 4 1
// Given an input k, return the kth row of Pascal's triangle.
// Bonus: Can you do this using only O(k) space?

func pascalTriangle(k int) []int {
	row := make([]int, k+1)
	row[0] = 1
	for r := 0; r < k; r++ {
		row[0], row[r+1] = 1, 1
		for i := r; i > 0; i-- {
			row[i] = row[i-1] + row[i]
		}
	}
	return row
}
