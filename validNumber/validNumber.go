// Source : https://oj.leetcode.com/problems/valid-number/
// Author : daoran
// Date   : 2016-04-09

/**********************************************************************************
*
* Validate if a given string is numeric.
*
* Some examples:
* "0" => true
* "   0.1  " => true
* "abc" => false
* "1 a" => false
* "2e10" => true
*
* Note: It is intended for the problem statement to be ambiguous.
*       You should gather all requirements up front before implementing one.
*
*
**********************************************************************************/

package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func isNumber(s string) bool {
	t := strings.TrimSpace(s)
	i, n := 0, len(t)

	if i < n && (string(t[i]) == "+" || string(t[i]) == "-") {
		i++
	}

	isNumeric := false
	for i < n && isDigit(t[i]) {
		i++
		isNumeric = true
	}

	if i < n && string(t[i]) == "." {
		i++
		for i < n && isDigit(t[i]) {
			i++
			isNumeric = true
		}
	}

	if isNumeric && i < n && string(t[i]) == "e" {
		i++
		isNumeric = false
		if i < n && (string(t[i]) == "+" || string(t[i]) == "-") {
			i++
		}
		for i < n && isDigit(t[i]) {
			i++
			isNumeric = true
		}
	}

	return isNumeric && i == n
}

func isDigit(d byte) bool {
	digit, _ := utf8.DecodeRune([]byte(string(d)))
	return unicode.IsDigit(digit)
}

func main() {
	fmt.Println(isNumber("1234"))
	fmt.Println(isNumber("."))
	fmt.Println(isNumber(" "))
	fmt.Println(isNumber("0e"))
	fmt.Println(isNumber("4e+"))
	fmt.Println(isNumber(".e1"))
	fmt.Println(isNumber(" -."))
}
