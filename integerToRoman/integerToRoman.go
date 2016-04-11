// Source : https://oj.leetcode.com/problems/integer-to-roman/
// Author : daoran
// Date   : 2016-04-11

/**********************************************************************************
*
* Given an integer, convert it to a roman numeral.
*
* Input is guaranteed to be within the range from 1 to 3999.
*
**********************************************************************************/

package main

import (
	"fmt"
	"strings"
)

var values = []int{1000, 900, 500, 400,
	100, 90, 50, 40,
	10, 9, 5, 4,
	1,
}

var symbols = []string{"M", "CM", "D", "CD",
	"C", "XC", "L", "XL",
	"X", "IX", "V", "IV",
	"I",
}

func intToRoman(num int) string {
	result := []string{}

	i := 0
	for num > 0 {
		k := num / values[i]
		for j := 0; j < k; j++ {
			result = append(result, symbols[i])
			num -= values[i]
		}
		i++
	}
	return strings.Join(result, "")
}

func main() {
	fmt.Println(intToRoman(1934))
	fmt.Println(intToRoman(500))
}
