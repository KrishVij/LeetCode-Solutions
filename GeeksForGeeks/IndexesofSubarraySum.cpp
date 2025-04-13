#include<vector>

std::vector<int> subarraySum(std::vector<int> &arr, int target)
{
    std::vector<int> indicesOfSum;
    int n = arr.size();
    int Initialptr {0};
    int Endptr {0};
    int sum {0};
    
    while (Endptr < n)
    {
        sum += arr[Endptr];
        while (sum > target)
        {
            sum -= arr[Initialptr];
            ++Initialptr;
        }
        if (sum == target)
        {
            indicesOfSum.emplace_back(Initialptr + 1);
            indicesOfSum.emplace_back(Endptr + 1);
        }
        ++Endptr;
    }
    return indicesOfSum;
}