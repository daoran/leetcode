// Source : https://oj.leetcode.com/problems/evaluate-reverse-polish-notation/
// Author : daoran
// Date   : 2016-04-11

/**********************************************************************************
*
* Evaluate the value of an arithmetic expression in Reverse Polish Notation.
*
* Valid operators are +, -, *, /. Each operand may be an integer or another expression.
*
* Some examples:
*
*   ["2", "1", "+", "3", "*"] -> ((2 + 1) * 3) -> 9
*   ["4", "13", "5", "/", "+"] -> (4 + (13 / 5)) -> 6
*
*
**********************************************************************************/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

var operators = "+-*/"

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		if strings.Index(operators, token) != -1 {
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			stack = append(stack, eval(x, y, token))
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func eval(x, y int, operator string) int {
	switch operator {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	default:
		return x / y
	}
}

func main() {
	fmt.Println(evalRPN([]string{"1", "2", "-"}))
}
