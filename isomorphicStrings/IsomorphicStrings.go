// Source : https://leetcode.com/problems/isomorphic-strings/
// Author : daoran
// Date   : 2016-04-14

/**********************************************************************************
 *
 * Given two strings s and t, determine if they are isomorphic.
 *
 * Two strings are isomorphic if the characters in s can be replaced to get t.
 *
 * All occurrences of a character must be replaced with another character while preserving
 * the order of characters. No two characters may map to the same character but a character
 *  may map to itself.
 *
 * For example,
 *
 *     Given "egg", "add", return true.
 *
 *     Given "foo", "bar", return false.
 *
 *     Given "paper", "title", return true.
 *
 * Note:
 * You may assume both s and t have the same length.
 *
 **********************************************************************************/

package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	maps := make([]int, 256)
	mapt := make([]int, 256)

	for i := 0; i < len(s); i++ {
		sInt := int(s[i])
		tInt := int(t[i])
		if maps[sInt] == 0 && mapt[tInt] == 0 {
			maps[sInt] = tInt
			mapt[tInt] = sInt
			continue
		}

		if maps[sInt] == tInt && mapt[tInt] == sInt {
			continue
		}

		return false
	}

	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
}
