#include<vector>
#include<algorithm>
int maxScore(std::vector<int>& cardPoints, int k)
{
    int n = cardPoints.size() - 1;
    int Initialptr {0};
    int Endptr {n};
    int sum {0};
    int maxSum {0};
    
    if (n == 0 || k <= 0)
    {
        return 0;
    }
    while (Initialptr <= k - 1)
    {
        sum = sum + cardPoints[Initialptr];
        ++Initialptr;
    }
    maxSum = sum;
    Initialptr = k - 1;
    while (Initialptr >= 0)
    {
        sum = sum - cardPoints[Initialptr];
        sum = sum + cardPoints[Endptr];
        maxSum = std::max(maxSum,sum);
        --Initialptr;
        --Endptr;
    }
    return maxSum;
}