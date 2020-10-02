package dcp

// Given a number in Roman numeral format, convert it to decimal.
// The values of Roman numerals are as follows:
// {'M': 1000, 'D': 500, 'C': 100, 'L': 50, 'X': 10, 'V': 5, 'I': 1}
// In addition, note that the Roman numeral system uses subtractive notation for numbers such as IV and XL.
// For the input XIV, for instance, you should return 14.

func romanToDecimal(roman string) int {
	m := map[rune]int{'M': 1000, 'D': 500, 'C': 100, 'L': 50, 'X': 10, 'V': 5, 'I': 1}
	decimal, current, length := 0, 1, len(roman)

	for i := length - 1; i >= 0; i-- {
		if n := m[rune(roman[i])]; n >= current {
			current = n
			decimal += n
		} else {
			decimal -= n
		}
	}

	return decimal
}
