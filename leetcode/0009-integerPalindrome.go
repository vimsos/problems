package leetcode

// Write a program that checks whether an integer is a palindrome. For example, 121 is a palindrome, as well as 888. 678 is not a palindrome. Do not convert the integer into a string.

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	number, reversed := x, 0
	for x/10 > 0 {
		reversed *= 10
		reversed += x % 10
		x /= 10
	}
	reversed *= 10
	reversed += x % 10
	return number == reversed
}
