#include<vector>

int maxSatisfied(std::vector<int>& customers, std::vector<int>& grumpy, int minutes)
{
    int n = customers.size();
    int g = grumpy.size();
    int Initialptr {0};
    int Endptr {0};
    int len {0};
    int maxSum {0};
    int defaultSum {0};
    int forcedSum {0};
    int totalSum {0};

    for (int i = 0;i<g && i<n;++i)
    {
        if (grumpy[i] == 0)
        {
            defaultSum += customers[i];
        }
    }

    while (Endptr < n)
    {
        if (grumpy[Endptr] == 1)
        {
            forcedSum += customers[Endptr];
        }

        len = Endptr - Initialptr + 1;
        if (len > minutes)
        {
            if (grumpy[Initialptr] == 1)
            {
                forcedSum -= customers[Initialptr];
            }
            
            ++Initialptr;
        }
        
        maxSum = std::max(maxSum,forcedSum);
        ++Endptr;

    }
    totalSum = defaultSum + maxSum;
    return maxSum;


}