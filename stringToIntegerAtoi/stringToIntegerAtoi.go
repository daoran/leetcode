package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var maxDiv10 = (1<<31 - 1) / 10

func myAtoi(str string) int {
	result := 0
	temp := strings.TrimSpace(str)
	i, n := 0, len(temp)
	sign := 1
	if i < n && string(temp[i]) == "+" {
		i++
	} else if i < n && string(temp[i]) == "-" {
		i++
		sign = -1
	}

	num := 0
	for i < n {
		r, _ := utf8.DecodeRune([]byte(string(temp[i])))
		if unicode.IsDigit(r) {
			digit, _ := strconv.Atoi(string(temp[i]))
			if num > maxDiv10 || (num == maxDiv10 && digit >= 8) {
				if sign == 1 {
					result = 1<<31 - 1
				} else {
					result = -1 << 31
				}
				return result
			}

			num = num*10 + digit
			i++
		} else {
			break
		}
	}

	result = sign * num
	return result
}

func main() {
	fmt.Println(myAtoi("   010"))
	fmt.Println(myAtoi("   -0012a42"))
	fmt.Println(myAtoi("-2147483648"))
	fmt.Println(myAtoi("2147483648"))
}
