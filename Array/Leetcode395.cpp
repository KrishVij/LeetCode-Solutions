#include<string>
#include<climits>
#include<algorithm>

int longestSubstring(std::string s, int k)
{
    int hash[256] = {0};
    int n = s.length();
    int uniquecounter {0};  // Count of unique characters
    int Initalptr {0};
    int Endptr {0};
    int len {0};
    int maxLen = INT_MIN;
    
    while (Endptr < n)
    {
        if (hash[s[Endptr]] == 0)
        {
            uniquecounter++;
        }
        hash[s[Endptr]]++;  //we increase the count of the element
        while (uniquecounter >= k)  //this conditions basically tells us that each element is unique and their frequency is greater than k due to earlier increase in the character count
        {
            len = Endptr - Initalptr + 1;
            maxLen = std::max (maxLen,len);
            hash[s[Initalptr]]--;  //as we shrink the window we decrease the count of the character that was at Initialptr and then we increment Initialptr
            if (hash[s[Initalptr]] == 0) 
            {
                uniquecounter--; // One less unique character
            }
            Initalptr++;

        }
        Endptr++;
    }
    return maxLen;
}
