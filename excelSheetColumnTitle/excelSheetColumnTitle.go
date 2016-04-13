// Source : https://oj.leetcode.com/problems/excel-sheet-column-title/
// Author : daoran
// Date   : 2016-04-13

/**********************************************************************************
 *
 * Given a positive integer, return its corresponding column title as appear in an Excel sheet.
 *
 * For example:
 *
 *     1 -> A
 *     2 -> B
 *     3 -> C
 *     ...
 *     26 -> Z
 *     27 -> AA
 *     28 -> AB
 *
 * Credits:Special thanks to @ifanchu for adding this problem and creating all test cases.
 *
 **********************************************************************************/

package main

import (
	"fmt"
	"strings"
)

func convertToTitle(n int) string {
	ret := []string{}

	for n > 0 {
		ch := 'A' + (n-1)%26
		ret = append([]string{fmt.Sprintf("%c", ch)}, ret[:]...)
		n -= (n - 1) % 26
		n /= 26
	}

	return strings.Join(ret, "")
}

func main() {
	fmt.Println(convertToTitle(3))
	fmt.Println(convertToTitle(26))
}
