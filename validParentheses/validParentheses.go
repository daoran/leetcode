// Source : https://oj.leetcode.com/problems/valid-parentheses/
// Author : daoran
// Date   : 2016-04-08

/**********************************************************************************
*
* Given a string containing just the characters '(', ')', '{', '}', '[' and ']',
* determine if the input string is valid.
*
* The brackets must close in the correct order, "()" and "()[]{}" are all valid
* but "(]" and "([)]" are not.
*
*
**********************************************************************************/

package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}

	for _, ch := range s {
		if ch == '{' || ch == '[' || ch == '(' {
			stack = append(stack, ch)
		} else if ch == '}' || ch == ']' || ch == ')' {
			if len(stack) == 0 {
				return false
			}

			sch := stack[len(stack)-1]
			if sch == '{' && ch == '}' || sch == '[' && ch == ']' || sch == '(' && ch == ')' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("{{}{[]()}}"))
}
