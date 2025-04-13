#include<vector>

int numSubarrayProductLessThanK(std::vector<int>& nums, int k)
{
    int n = nums.size();
    int Initialptr {0};
    int Endptr {0};
    int subarrCount = nums.size();
    int product {1};

    if (nums.empty() || k == 0)
    {
        return 0;
    }
    while (Endptr < n) 
    {
        product *= nums[Endptr];

        while (product >= k && Initialptr <= Endptr) 
        {
            product /= nums[Initialptr];
            ++Initialptr;
        }
        subarrCount += Endptr - Initialptr + 1;
        ++Endptr;
    }
    return subarrCount;
}