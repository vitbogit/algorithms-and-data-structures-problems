# Concise solution

The idea is simple. For each character, remember its last appearance index using the hash map. Then increase the substring and check if the last occurrence of the new element is within that substring. If yes, compare the length of the substring with the best length and start a new substring immediately after the last occurrence of the new character (formula for new substring length: i - LastAppearance).

One more thing: by the end of the string, there may be characters left in your "current substring" counter that you may not have compared to the "longest substring" counter.