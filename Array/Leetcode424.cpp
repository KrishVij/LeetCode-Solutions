#include<vector>
#include<string>
#include<unordered_map>
#include<algorithm>

int characterReplacement(std::string s, int k)
{
    std::vector<int>alpha (26,0);
    int n = s.length();
    int Initialptr {0};
    int Endptr {0};
    int len {0};
    int maxLen {0};
    int maxF {0};

    while (Endptr < n)
    {
        len = Endptr - Initialptr + 1;
        alpha[s[Endptr] - 'A']++;
        maxF = std::max(maxF,alpha[s[Endptr] - 'A']);
        if (len - maxF > k)
        {
            alpha[s[Initialptr] - 'A']--;
            maxF = 0;
            ++Initialptr;
        }
        if (len - maxF <= k)
        {
            maxLen = std::max(maxLen,len);
        }
        ++Endptr;
    }
    return maxLen;


}

