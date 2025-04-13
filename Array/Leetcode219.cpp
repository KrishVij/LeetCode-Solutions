#include<vector>
#include<unordered_map>
#include<algorithm>

bool containsNearbyDuplicate(std::vector<int>& nums, int k)
{
    std::unordered_map<int,int> num_map;
    int n = nums.size();
    for (int i = 0;i<n;++i)
    {
        if (num_map.count(nums[i]))
        {
            if (std::abs(i - num_map[nums[i]]))
            {
                return true;
                break;
            }
        }
    }
    return false;
}
