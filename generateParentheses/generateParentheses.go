// Source : https://oj.leetcode.com/problems/generate-parentheses/
// Author : daoran
// Date   : 2016-04-16

/**********************************************************************************
*
* Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
*
* For example, given n = 3, a solution set is:
*
* "((()))", "(()())", "(())()", "()(())", "()()()"
*
*
**********************************************************************************/

package main

import "fmt"

func generateParenthesis(n int) []string {
	result := &[]string{}

	var s string
	generator(result, n, n, s)
	return *result
}

func generator(result *[]string, left int, right int, s string) {
	if left == 0 && right == 0 {
		*result = append(*result, s)
		return
	}

	if left > 0 {
		generator(result, left-1, right, s+"(")
	}

	if right > 0 && right > left {
		generator(result, left, right-1, s+")")
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
