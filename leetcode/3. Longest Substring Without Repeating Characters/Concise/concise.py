class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        d = {}
        cur, best = 0, 0
        for i in range(len(s)):
            if s[i] in d and d[s[i]] >= i-cur:
                best = max(best, cur)
                cur = i-d[s[i]]
            else:   
                cur += 1
            d[s[i]] = i
        return max(best, cur)