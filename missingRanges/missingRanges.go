// Source : https://oj.leetcode.com/problems/missing-ranges/
// Author : daoran
// Date   : 2016-04-09

/**********************************************************************************
 *
 * Given a sorted integer array where the range of elements are [lower, upper] inclusive,
 * return its missing ranges.
 *
 * For example, given [0, 1, 3, 50, 75], lower = 0 and upper = 99,
 * return ["2", "4->49", "51->74", "76->99"].
 *
 **********************************************************************************/

package main

import (
	"fmt"
	"strconv"
)

func findMissingRanges(vals []int, start, end int) []string {
	ranges := []string{}
	prev := start - 1
	for i := 0; i <= len(vals); i++ {
		curr := 0
		if i == len(vals) {
			curr = end + 1
		} else {
			curr = vals[i]
		}

		if (curr - prev) >= 2 {
			ranges = append(ranges, getRange(prev+1, curr-1))
		}

		prev = curr
	}

	return ranges
}

func getRange(from, to int) string {
	if from == to {
		return strconv.Itoa(from)
	}

	return strconv.Itoa(from) + "->" + strconv.Itoa(to)
}

func main() {
	fmt.Println(findMissingRanges([]int{0, 1, 3, 50, 75}, 0, 99))
}
