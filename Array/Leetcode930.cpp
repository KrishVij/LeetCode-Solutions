#include<vector>

int numSubarraysWithSum(std::vector<int>& nums, int goal)
{
    int n = nums.size();
    int Initialptr {0};
    int Endptr {0};
    int len {0};
    int sum {0};
    int count {0};

    if (goal < 0)
    {
        return 0;
    }

    while (Endptr <= n)
    {
        sum += nums[Endptr];
        while (sum > goal)
        {
            sum -= nums[Initialptr];
            ++Initialptr;
        }
        len = Endptr - Initialptr + 1;
        count += len;
        ++Endptr;
    }
    int result = count - numSubarraysWithSum(nums, goal - 1);
    return result;
}