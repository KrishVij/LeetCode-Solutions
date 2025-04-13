#include<vector>
#include<algorithm>
#include<climits>

std::vector<int> max_of_subarrays(int k, std::vector<int> &arr)
{
    std::vector<int> maxContainer;
    int n = arr.size();
    int len {0};
    int Initialptr {0};
    int Endptr {0};
    int maxEle {0};

    while (Endptr < n)
    {
        len = Endptr - Initialptr + 1;
        if (len < k)
        {
            ++Endptr;
        }
        else if (len == k)
        {
            maxEle = INT_MIN;
            for (int i = Initialptr;i<Endptr + 1;++i)
            {
                maxEle = std::max(maxEle,arr[i]);
            }
        }
        maxContainer.emplace_back(maxEle);
        
        ++Initialptr;
        ++Endptr;
    }
    return maxContainer;

}
