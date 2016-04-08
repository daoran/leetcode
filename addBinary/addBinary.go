// Source : https://oj.leetcode.com/problems/add-binary/
// Author : daoran
// Date   : 2016-04-08

/**********************************************************************************
*
* Given two binary strings, return their sum (also a binary string).
*
* For example,
* a = "11"
* b = "1"
* Return "100".
*
*
**********************************************************************************/

package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	alen := len(a)
	blen := len(b)
	carry := 0
	var result string

	for alen > 0 || blen > 0 {
		abit := 0
		bbit := 0
		if alen <= 0 {
			abit = 0
		} else {
			abit, _ = strconv.Atoi(string(a[alen-1]))
		}

		if blen <= 0 {
			bbit = 0
		} else {
			bbit, _ = strconv.Atoi(string(b[blen-1]))
		}

		val := (abit + bbit + carry) % 2
		carry = (abit + bbit + carry) / 2
		result = strconv.Itoa(val) + result
		alen--
		blen--
	}
	if carry == 1 {
		result = "1" + result
	}
	return result
}

func main() {
	fmt.Println(addBinary("10", "1"))
}
