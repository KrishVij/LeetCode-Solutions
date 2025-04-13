#include<vector>
#include<string>
#include<unordered_map>
#include<climits>

std::string minWindow(std::string s, std::string t)
{
    std::unordered_map<char,int>hash;
    int n = s.length();
    int m = t.length();
    int Initialptr {0};
    int Endptr {0};
    int startIndex {-1};
    int len {0};
    int minLen {INT_MAX};
    int count {0};
    
    for (int i = 0;i<m;++i)
        {
            hash[t[i]]++;
        }

    // int requiredMatches = hash.size();

    while (Endptr < n)
    {
        len = Endptr - Initialptr + 1;
        if (hash[s[Endptr]] > 0)
        {
            count++;
        }
        else
        {
            hash[s[Endptr]]--;
        }
        while (count == m)
        {
            if (len < minLen)
            {
                minLen = len;
                startIndex = Initialptr;
                
            }
            hash[s[Initialptr]]--;
            ++Initialptr;
            if (hash[s[Initialptr]] > 0)
            {
                count --;
            }
        }
        ++Endptr;
    }
    return startIndex == -1 ? "" : s.substr(startIndex,minLen);
}