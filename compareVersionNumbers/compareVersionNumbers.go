// Source : https://oj.leetcode.com/problems/compare-version-numbers/
// Author : daoran
// Date   : 2016-04-08

/**********************************************************************************
 *
 * Compare two version numbers version1 and version1.
 * If version1 > version2 return 1, if version1 < version2 return -1, otherwise return 0.
 *
 * You may assume that the version strings are non-empty and contain only digits and the . character.
 * The . character does not represent a decimal point and is used to separate number sequences.
 * For instance, 2.5 is not "two and a half" or "half way to version three", it is the fifth second-level revision of the second first-level revision.
 *
 * Here is an example of version numbers ordering:
 * 0.1 < 1.1 < 1.2 < 13.37
 *
 * Credits:Special thanks to @ts for adding this problem and creating all test cases.
 *
 **********************************************************************************/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	// pad 0 case 1.0.1 1
	for len(v1) < len(v2) {
		v1 = append(v1, "0")

	}

	for len(v1) > len(v2) {
		v2 = append(v2, "0")
	}

	for i := 0; i < len(v1); i++ {
		// case 01 1
		v1, _ := strconv.Atoi(v1[i])
		v2, _ := strconv.Atoi(v2[i])

		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
	}

	return 0
}

func main() {
	fmt.Println(compareVersion("1.1", "1.10"))
	fmt.Println(compareVersion("1.2", "1.10"))
}
