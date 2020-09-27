package dcp

// Spreadsheets often use this alphabetical encoding for its columns: "A", "B", "C", ..., "AA", "AB", ..., "ZZ", "AAA", "AAB", ....
// Given a column number, return its alphabetical column id. For example, given 1, return "A". Given 27, return "AA".

func columnIDToAlpha(id int) string {
	var alpha string
	for id > 0 {
		r := id % 26
		id /= 26
		if r == 0 {
			r, id = 26, id-1
		}
		alpha = string('A'+r-1) + alpha
	}
	return alpha
}
