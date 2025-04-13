#include<vector>

int atMostKOddNumbers(const std::vector<int>& nums, int k) 
{

        int n = nums.size();
        int Initialptr{0};
        int Endptr{0};
        int len{0};
        int sum{0};
        int count{0};

        if (k < 0) 
        {
            return 0;
        }

        while (Endptr < n) 
        {
            sum += nums[Endptr] % 2;
            while (sum > k) 
            {
                sum -= nums[Initialptr] % 2;
                ++Initialptr;
            }
            len = Endptr - Initialptr + 1;
            count += len;
            ++Endptr;
        }
        return count;
}

    int numberOfSubarrays(std::vector<int>& nums, int k) 
    {
        return atMostKOddNumbers(nums, k) - atMostKOddNumbers(nums, k - 1);
    }