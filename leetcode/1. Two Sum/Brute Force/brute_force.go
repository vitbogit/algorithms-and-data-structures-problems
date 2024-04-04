func twoSum(nums []int, target int) []int {
	for i := range nums {
		for j := 1; j < len(nums); j++ {
			if i != j && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
