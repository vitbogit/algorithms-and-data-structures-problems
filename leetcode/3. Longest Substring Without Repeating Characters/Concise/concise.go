func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	cur, best := 0, 0
	for i := 0; i < len(s); i++ {
		if lastAppearence, ok := m[s[i]]; ok && lastAppearence >= i-cur {
			best = max(best, cur)
			cur = i - lastAppearence
		} else {
			cur += 1
		}
		m[s[i]] = i
	}
	return max(best, cur)
}