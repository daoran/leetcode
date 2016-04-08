// Source : https://leetcode.com/problems/integer-to-english-words/
// Author : Hao Chen
// Date   : 2015-10-22

/***************************************************************************************
 *
 * Convert a non-negative integer to its english words representation. Given input is
 * guaranteed to be less than 231 - 1.
 *
 * For example,
 *
 * 123 -> "One Hundred Twenty Three"
 * 12345 -> "Twelve Thousand Three Hundred Forty Five"
 * 1234567 -> "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
 *
 *   Did you see a pattern in dividing the number into chunk of words? For example, 123
 * and 123000.
 *
 *   Group the number by thousands (3 digits). You can write a helper function that
 * takes a number less than 1000 and convert just that chunk to words.
 *
 *   There are many edge cases. What are some good test cases? Does your code work with
 * input such as 0? Or 1000010? (middle chunk is zero and should not be printed out)
 *
 ***************************************************************************************/

package main

import "fmt"

var dict1 = []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine",
	"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
var dict2 = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

var dict3 = []string{"Hundred", "Thousand", "Million", "Billion"}

func numberLess1000ToWords(num int) string {
	var result string

	if num == 0 {
		return result
	} else if num < 20 {
		return dict1[num]
	} else if num < 100 {
		result = dict2[num/10]
		if num%10 > 0 {
			result += " " + dict1[num%10]
		}
	} else {
		result = dict1[num/100] + " " + dict3[0]
		if num%100 > 0 {
			result += " " + numberLess1000ToWords(num%100)
		}
	}
	return result
}

func numberToWords(num int) string {
	if num == 0 {
		return dict1[num]
	}

	ret := []string{}
	for ; num > 0; num /= 1000 {
		ret = append(ret, numberLess1000ToWords(num%1000))
	}

	result := ret[0]
	for i := 1; i < len(ret); i++ {
		if len(ret[i]) > 0 {
			if len(result) > 0 {
				result = ret[i] + " " + dict3[i] + " " + result
			} else {
				result = ret[i] + " " + dict3[i]
			}
		}
	}

	return result
}

func main() {
	fmt.Println(numberToWords(12345))
}
