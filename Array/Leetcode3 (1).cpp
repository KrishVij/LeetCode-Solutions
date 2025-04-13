#include<vector>
#include<algorithm>
#include<string>

int lengthOfLongestSubstring(std::string s)
{
    std::vector<int> hash (256,-1);
    int n = s.length();
    int initialptr = 0;
    int endptr = 0;
    int len = 0;
    int maxlen = 0;

    while (endptr < n)
    {
        if (hash[s[endptr]] != -1)
        {
            if (hash[s[endptr]] >= initialptr)
            {
                initialptr = hash[s[endptr]] + 1;

            }
        }
        len = endptr - initialptr + 1;
        maxlen = std::max(maxlen,len);
        hash[s[endptr]] = endptr;
        ++endptr;
    }
    return maxlen;


}
 //Time complesxity : O(n);
 //Space complexity : O(256) -> as we made a hash array of 256;