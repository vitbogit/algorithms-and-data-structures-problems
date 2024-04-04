class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        for(int i=0; i<nums.size(); i++){
            for(int j=0; j<nums.size(); j++){
                if(i!=j&&target==nums[i]+nums[j]){
                    return vector<int>{i,j};
                }
            }
        }
        return vector<int>{};
    }
};