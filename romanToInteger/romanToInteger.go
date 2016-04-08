// Source : https://oj.leetcode.com/problems/roman-to-integer/
// Author : daoran
// Date   : 2016-04-08

/**********************************************************************************
*
* Given a roman numeral, convert it to an integer.
*
* Input is guaranteed to be within the range from 1 to 3999.
*
**********************************************************************************/

package main

import "fmt"

func romanCharToInt(char string) int {
	d := 0
	switch char {
	case "I":
		d = 1
	case "V":
		d = 5
	case "X":
		d = 10
	case "L":
		d = 50
	case "C":
		d = 100
	case "D":
		d = 500
	case "M":
		d = 1000
	}

	return d
}

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := romanCharToInt(string(s[0]))
	for i := 1; i < len(s); i++ {
		prev := romanCharToInt(string(s[i-1]))
		curr := romanCharToInt(string(s[i]))
		if prev < curr {
			result = result - prev + curr - prev
		} else {
			result += curr
		}
	}
	return result
}

func main() {
	fmt.Println(romanToInt("MCMIV"))
}
