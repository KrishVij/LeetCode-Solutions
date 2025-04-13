#include<vector>
#include<unordered_set>

 bool subArrayExists(std::vector<int> arr) 
    {
        std::unordered_set<int> prefixSums;
        int sum = 0;
        
        for (int i : arr)
        {
            sum += i;
            if (sum == 0 || prefixSums.find(sum) != prefixSums.end()) 
            {
            return true;
            }
            prefixSums.insert(sum);
        }
        return false;
    }