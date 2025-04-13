#include <climits>
#include <vector>
int maxProfit(std::vector<int> &prices) 
{
  int index = 0;
  int buyday = prices[0];
  int sellday = prices[0];
  int profit = 0;
  while (index<prices.size() - 1)
  {
    while (prices[index] > prices[index+1])
    {
      ++index;
    }
      buyday = prices[index];
 
    while (prices[index]<prices[index+1])
    {
      ++index;
    }
    sellday = prices[index];
    profit += sellday - buyday;
 }
  return profit;
}
 
  

