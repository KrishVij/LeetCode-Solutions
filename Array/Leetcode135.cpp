#include <algorithm>
#include <vector>
int candy(std::vector<int> &ratings) {
  std::vector<int> candies(ratings.size(), 1);
  int sum = 0;
  for (int left = 1; left < candies.size(); ++left) {
    if (ratings[left] > ratings[left - 1]) {
      candies[left] = candies[left + 1] + 1;
    }
  }

  for (int right = candies.size() - 1; right >= 0; --right) {
    if (ratings[right] > ratings[right - 1]) {
      candies[right] = std::max(candies[right], candies[right + 1] + 1);
    }
  }
  for (int i = 0; i < candies.size(); ++i) {
    sum += candies[i];
  }
  return sum;
}
