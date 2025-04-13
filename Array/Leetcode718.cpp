#include<vector>
#include<algorithm>

int findLength(std::vector<int>& nums1, std::vector<int>& nums2)
{
    int Initialptr {0};
    int n1 = nums1.size();
    int n2 = nums2.size();
    int Endptr {0};
    int len {0};
    int maxLen {0};

    while (Endptr < n1 && Endptr < n2)
    {
        
        while (nums1[Endptr] == nums2[Endptr])
        {
            len = Endptr - Initialptr + 1;
            maxLen = std::max(maxLen,len);
            ++Endptr;
        }
        if (nums1[Endptr] != nums2[Endptr])
        {
            ++Initialptr;
        }
        ++Endptr;
    }
    return maxLen;
}

// 21 out of 55 test cases passed