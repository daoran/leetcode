// Source : https://oj.leetcode.com/problems/zigzag-conversion/
// Author : Hao Chen
// Date   : 2014-07-17

/**********************************************************************************
*
* The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
* (you may want to display this pattern in a fixed font for better legibility)
*
* P   A   H   N
* A P L S I I G
* Y   I   R
*
* And then read line by line: "PAHNAPLSIIGYIR"
*
* Write the code that will take a string and make this conversion given a number of rows:
*
* string convert(string text, int nRows);
*
* convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".
*
*
**********************************************************************************/

package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows <= 1 || numRows >= len(s) {
		return s
	}

	r := make([]string, numRows)
	row := 0
	step := 1
	for i := 0; i < len(s); i++ {
		if row == numRows-1 {
			step = -1
		}
		if row == 0 {
			step = 1
		}

		r[row] += string(s[i])
		row += step
	}

	var result string
	for i := 0; i < numRows; i++ {
		result += r[i]
	}

	return result
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
