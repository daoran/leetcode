// Source : https://oj.leetcode.com/problems/one-edit-distance/
// Author : daoran
// Date   : 2016-04-10

/*
 *    Given two strings S and T, determine if they are both one edit distance apart.
 */

package main

import "fmt"

func isOneEditDistance(s, t string) bool {
	m := len(s)
	n := len(t)
	if m > n {
		return isOneEditDistance(t, s)
	}

	if n-m > 1 {
		return false
	}

	i := 0
	shift := n - m

	for i < m && string(s[i]) == string(t[i]) {
		i++
	}

	if i == m {

		return shift > 0
	}

	if shift == 0 {
		i++
	}

	for i < m && string(s[i]) == string(t[i+shift]) {
		i++
	}

	return i == m
}

func main() {
	fmt.Println(isOneEditDistance("abcde", "abXde"))
	fmt.Println(isOneEditDistance("abcde", "abcXde"))
	fmt.Println(isOneEditDistance("abcde", "abcdeeX"))
}
