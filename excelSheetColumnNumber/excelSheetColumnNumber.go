// Source : https://oj.leetcode.com/problems/excel-sheet-column-number/
// Author : daoran
// Date   : 2016-04-12

/**********************************************************************************
 *
 * Related to question Excel Sheet Column Title
 * Given a column title as appear in an Excel sheet, return its corresponding column number.
 *
 * For example:
 *     A -> 1
 *     B -> 2
 *     C -> 3
 *     ...
 *     Z -> 26
 *     AA -> 27
 *     AB -> 28
 *
 * Credits:Special thanks to @ts for adding this problem and creating all test cases.
 *
 **********************************************************************************/

package main

import "fmt"

func titleToNumber(s string) int {
	ret := 0
	for _, r := range s {
		n := int(r-'A') + 1
		ret = ret*26 + n
	}

	return ret
}

func main() {
	fmt.Println(titleToNumber("AB"))
}
