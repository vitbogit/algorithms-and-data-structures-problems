# Hash Map, One Pass solution

Use a hash map where the keys represent the numbers in the array nums and the values represent the indices of those numbers in that array.

Make a pass. For each element in nums, check whether the corresponding target-num[i] element is already present in the hash map. If yes, return the answer. Otherwise, add num[i] to the hash map.

Time complexity - O(N)
Space complexity - O(N)