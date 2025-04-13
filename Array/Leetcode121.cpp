#include <limits>
#include <vector>
int maxProfit(std::vector<int> &prices) {
  int maximum = std::numeric_limits<int>::min();
  int minimum = std::numeric_limits<int>::max();
  if (prices.empty()) {
    return 0;
  }
  for (int i = 0; i < prices.size(); ++i) {
    if (prices[i] < minimum) {
      minimum = prices[i];
    }
    int profit = prices[i] - minimum;
    if (profit > maximum) {
      maximum = profit;
    }
  }

  return maximum;
}
