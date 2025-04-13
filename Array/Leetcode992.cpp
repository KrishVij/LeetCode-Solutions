#include<vector>
#include<unordered_map>

int subarraysAtleastKDistinct(std::vector<int>& nums, int k)
{
    std::unordered_map<int,int> elementStore;
    int n = nums.size();
    int Initialptr {0};
    int Endptr {0};
    int count {0};
    int len {0};

    while (Endptr < n)
    {
        nums[elementStore[Endptr]]++;
        while (elementStore.size() > k)
        {
            nums[elementStore[Initialptr]]--;
            if (nums[elementStore[Initialptr]] == 0)
            {
                elementStore.erase(nums[Initialptr]);
            }
            ++Initialptr;
        }
        len = Endptr - Initialptr + 1;
        count += len;
        ++Endptr;
    }
    return count;
}
int subarraysWithKDistinct(std::vector<int>& nums, int k)
{
    int result = subarraysAtleastKDistinct(nums, k) - subarraysAtleastKDistinct(nums, k - 1);
    return result;
}