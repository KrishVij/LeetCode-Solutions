#include<vector>
#include<algorithm>
#include<iostream>


/*
BruteForce Soln:

int longestOnes(std::vector<int>& nums, int k)
{
    int zeroCounter;
    int len {0};
    int maxlen {0};

    for (int i = 0;i<nums.size();++i)
    {
        zeroCounter = 0;
        for (int j = i;j<nums.size();++j)
        {
            if (nums[j] == 0)
            {
                zeroCounter++;
            }
            else if (zeroCounter <= k)
            {
                len = j - i + 1;
                maxlen = std::max(maxlen,len);
                
            }
            else
            {
                break;
            }
            
        }
    }
    return maxlen;
}

TimeComplexity : O(n^2);
SpaceComplexity : O(1);
*/

int longestOnes(std::vector<int>& nums, int k)
{
    int n = nums.size();
    int Initalptr  {0};
    int Endptr {0};
    int zeroCounter {0};
    int len {0};
    int maxlen {0};

    while (Endptr < n)
    {
        if (nums[Endptr] == 0)
        {
            zeroCounter++;
        }
        while (zeroCounter > k)
        {
            if (nums[Initalptr] == 0)
            {
                zeroCounter--;
            }
            ++Initalptr;
        }
        if (zeroCounter <= k)
        {
            len = Endptr - Initalptr + 1;
            maxlen = std::max(maxlen,len);
            
            
        }
        ++Endptr;
    }
    return maxlen;
}




