package main

import "fmt"

func twoSum(nums []int, target int) []int {
    i, j := 0, len(nums) - 1
    result := []int{}
    for i < j {
        sum := nums[i] + nums[j]
        if sum < target {
            i++
        } else if sum > target {
            j--
        } else {
            result = []int{i, j}
            break
        }
    }
    
    return result 
}

func main() {
    fmt.Println(twoSum([]int{1,2,3,4,5}, 8))
}