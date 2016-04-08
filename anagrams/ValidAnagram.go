// Source : https://leetcode.com/problems/valid-anagram/
// Author : daoran
// Date   : 2016-04-08

/**********************************************************************************
 *
 * Given two strings s and t, write a function to determine if t is an anagram of s.
 *
 * For example,
 * s = "anagram", t = "nagaram", return true.
 * s = "rat", t = "car", return false.
 *
 * Note:
 * You may assume the string contains only lowercase alphabets.
 *
 **********************************************************************************/

package main

import "fmt"

func isAnagram(s string, t string) bool {
	chs := make(map[rune]int)
	for _, r := range s {
		chs[r]++
	}

	for _, r := range t {
		chs[r]--
	}

	for _, value := range chs {
		if value != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("rat", "cat"))
	fmt.Println(isAnagram("anagram", "nagaram"))
}
