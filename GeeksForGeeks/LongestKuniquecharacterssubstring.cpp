#include<string>
#include<algorithm>
#include<unordered_map>

int longestKSubstr(std::string s, int k) 
    {
        std::unordered_map<char,int> mpp;
        int n = s.length();
        int Initialptr = 0;
        int Endptr = 0;
        int len = 0;
        int maxLen = -1;
        
        while (Endptr < n)
        {
            mpp[s[Endptr]]++;
            while (mpp.size() > k)
            {
                mpp[s[Initialptr]]--;
                if (mpp[s[Initialptr]] == 0)
                {
                    mpp.erase(s[Initialptr]);
                }
                Initialptr++;
                
                
            }
            if (mpp.size() == k)
            {
                len = Endptr - Initialptr + 1;
                maxLen = std::max(maxLen,len);
                
            }
            Endptr++;
        }
        return maxLen;
    }

