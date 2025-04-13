#include<vector>
#include<algorithm>

int minSubArrayLen(int target, std::vector<int>& nums)
{
    int n = nums.size();
    int Initialptr {0};
    int Endptr {0};
    int len {0};
    int minLen = __INT_MAX__;
    int sum {0};

    while (Endptr < n)
    {
        sum += nums[Endptr];
        while (sum >= target)
        {
            sum -= nums[Initialptr];
            len = Endptr - Initialptr + 1;
            minLen = std::min(minLen,len);
            Initialptr++;
        }
        Endptr++;
    }
    if (minLen == __INT_MAX__)
    {
        return 0;
    }
    else
    {
        return minLen;
    }
}