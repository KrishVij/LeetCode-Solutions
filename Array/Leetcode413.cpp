#include<vector>

int numberOfArithmeticSlices(std::vector<int>& nums)
{
    int n = nums.size();
    int Initialptr {0};
    int Endptr {0};
    int len {0};
    int count {0};

    while (Endptr < n)
    {
        len = Endptr - Initialptr + 1;
        if (len >=3)
        {
            int diff1 = nums[Initialptr + 1] - nums[Initialptr];
            int diff2 = nums[Endptr] - nums[Endptr - 1];
            if (diff1 == diff2)
          {
            count++;
          }
        else
        {
            Initialptr = Endptr - 1;
        }
        }
        
        
        ++Endptr;
    }
    return count;
}