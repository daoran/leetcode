func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	result := []int{}
	for i := 0; i < len(nums); i++ {
		if j, ok := hash[target-nums[i]]; ok {
			result = []int{j, i}
			break
		}
		hash[nums[i]] = i
	}

	return result
}
