// Source : https://oj.leetcode.com/problems/climbing-stairs/
// Author : daoran
// Date   : 2016-04-12

/**********************************************************************************
*
* You are climbing a stair case. It takes n steps to reach to the top.
*
* Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
*
*
**********************************************************************************/

package main

import "fmt"

func climbStairs(n int) int {
	p := 1
	q := 1

	for i := 2; i <= n; i++ {
		temp := q
		q += p
		p = temp
	}

	return q
}

func main() {
	fmt.Println(climbStairs(3))
}
