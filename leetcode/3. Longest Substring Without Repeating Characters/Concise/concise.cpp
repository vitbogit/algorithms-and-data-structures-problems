class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        unordered_map <char, int> m;
        int cur = 0;
        int best = 0;
        for(int i=0; i<s.length(); i++){
            if(m.count(s[i])>0 && m[s[i]]>=i-cur){
                if(cur>best){
                    best = cur;
                }
                cur = i - m[s[i]];
            } else {
                cur += 1;
            }
            m[s[i]] = i;
        }
        if(cur>best){
            best = cur;
        }
        return best;
    }
};