func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if secondIndex, ok := m[complement]; ok {
			return []int{i, secondIndex}
		}
		m[num] = i
	}
	return []int{}
}