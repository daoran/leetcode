// Source : https://oj.leetcode.com/problems/count-and-say/
// Author : daoran
// Date   : 2016-04-08

/**********************************************************************************
*
* The count-and-say sequence is the sequence of integers beginning as follows:
* 1, 11, 21, 1211, 111221, ...
*
* 1 is read off as "one 1" or 11.
* 11 is read off as "two 1s" or 21.
* 21 is read off as "one 2, then one 1" or 1211.
*
* Given an integer n, generate the nth sequence.
*
* Note: The sequence of integers will be represented as a string.
*
*
**********************************************************************************/

package main

import (
	"fmt"
	"strconv"
)

func getNext(slice []int) []int {
	cnt := 0
	val := 0

	result := []int{}
	for i := 0; i < len(slice); i++ {
		if i == 0 {
			val = slice[0]
			cnt = 1
			continue
		}

		if slice[i] == val {
			cnt++
		} else {
			result = append(result, cnt)
			result = append(result, val)
			val = slice[i]
			cnt = 1
		}

	}

	if cnt > 0 && val > 0 {
		result = append(result, cnt)
		result = append(result, val)
	}

	return result

}

func countAndSay(n int) string {
	if n <= 0 {
		return ""
	}
	if n == 1 {
		return "1"
	}

	var result string
	slice := []int{}
	slice = append(slice, 1)

	for i := 2; i <= n; i++ {
		slice = getNext(slice)
	}

	for i := 0; i < len(slice); i++ {
		result += strconv.Itoa(slice[i])
	}

	return result
}

func main() {
	fmt.Println(countAndSay(5))
}
