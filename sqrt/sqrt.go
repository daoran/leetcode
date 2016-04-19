// Source : https://oj.leetcode.com/problems/sqrtx/
// Author : daoran
// Date   : 2016-04-19

/**********************************************************************************
*
* Implement int sqrt(int x).
*
* Compute and return the square root of x.
*
**********************************************************************************/

package main

import "fmt"

// http://en.wikipedia.org/wiki/Newton%27s_method
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	var last float64
	var ret float64 = 1

	for ret != last {
		last = ret
		ret = (ret + float64(x)/ret) / 2
	}

	return int(ret)
}

func main() {
	fmt.Println(mySqrt(81))
	fmt.Println(mySqrt(2147395599))
}
