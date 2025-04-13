#include<vector>

int maxTurbulenceSize(std::vector<int>& arr)
{
    int n = arr.size();
    int lastComparison {0};
    int Endptr {0};
    int turbLen {1};
    int maxLen {1};
    int comparison;

    while (Endptr < n)
    {
        if (arr[Endptr] > arr[Endptr - 1])
        {
            comparison = 1;
        }
        else if (arr[Endptr] < arr[Endptr - 1])
        {
            comparison = -1;
        }
        else
        {
            comparison = 0;
        }
        if (comparison == -lastComparison)
        {
            turbLen++;
        }
        else if (comparison == 0)
        {
            turbLen = 1;
        }
        else
        {
            turbLen = 2;
        }
        lastComparison = comparison;
        maxLen = std::max(maxLen, turbLen);
        Endptr++;
    }
    return maxLen;
}