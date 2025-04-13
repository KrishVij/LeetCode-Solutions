#include<string>
#include<algorithm>

int numberOfSubstrings(std::string s) 
{
        int LastSeen[3] = {-1, -1, -1};
        int n = s.length();
        int count = 0;

        for (int i = 0; i < n; ++i) 
        {
            LastSeen[s[i] - 'a'] = i;
            if (LastSeen[0] != -1 && LastSeen[1] != -1 && LastSeen[2] != -1) 
            {
                count += 1 + std::min({LastSeen[0], LastSeen[1], LastSeen[2]});
            }
        }
        return count;
}
