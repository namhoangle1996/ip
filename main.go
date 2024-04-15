package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	// Initialize a 2D dynamic programming table to store results of subproblems
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	// Initialize variables to keep track of the longest palindrome substring
	start, maxLen := 0, 1

	// Base case: single characters are palindromes
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	// Check for palindromes of length 2
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLen = 2
		}
	}

	// Check for palindromes of length 3 or greater
	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				start = i
				maxLen = length
			}
		}
	}

	// Return the longest palindrome substring
	return s[start : start+maxLen]
}

func main() {
	input := "babadaa"
	longest := longestPalindrome(input)
	fmt.Println("Longest palindrome substring:", longest)
}
