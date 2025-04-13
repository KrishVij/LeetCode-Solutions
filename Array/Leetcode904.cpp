#include<vector>
#include<unordered_map>
/*
int totalFruit(std::vector<int>& fruits)
{
    std::unordered_map<int,int> map;
    int n = fruits.size();
    int Initialptr {0};
    int Endptr {0};
    int len {0};
    int maxLen {0};

    while (Endptr < n)
    {
        map[fruits[Endptr]]++;
        if (map.size() > 2)
        {
            while (map.size() > 2)
            {
                map[fruits[Initialptr]]--;
                if (map[fruits[Initialptr]] == 0)
                {
                    map.erase(fruits[Initialptr]);
                }
                ++Initialptr;
            }
        }
        if (map.size() <= 2)
        {
            len = Endptr - Initialptr + 1;
            maxLen = std::max(maxLen,len);
        }
        ++Endptr;
    }

}
Time Complexity : O(N) + 0(N);
Space Complexity : O(3);
*/

//Most Optimal Solution

int totalFruit(std::vector<int>& fruits)
{
    std::unordered_map<int,int> map;
    int n = fruits.size();
    int Initialptr {0};
    int Endptr {0};
    int len  {0};
    int maxLen {0};

    while (Endptr < n)
    {
        map[fruits[Endptr]]++;
        if (map.size() > 2)
        {
            map[fruits[Initialptr]]--;
            if (map[fruits[Initialptr]] == 0)
            {
                map.erase(fruits[Initialptr]);
                ++Initialptr;
            }
        }
        if (map.size() < 2)
        {
            len = Endptr - Initialptr + 1;
            maxLen = std::max(maxLen,len);
        }
        ++Endptr;
    }
    return maxLen;
}

//Time Complexity : O(N);
//Space Complexity : O(3);