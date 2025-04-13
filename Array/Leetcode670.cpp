#include <algorithm>
#include <functional>
#include <string>
#include <vector>
int maximumSwap(int num) {
  std::vector<int> greatNumStore;
  std::string empty = "";
  int count = 0;
  if (num <= 0) {
    return num;
  }
  while (num > 0) {
    int digit = num % 10;
    ++count;
    greatNumStore.emplace_back(digit);
    num = num / 10;
  }
  std::reverse(greatNumStore.begin(), greatNumStore.end());
  if (greatNumStore.size() <= 1) {
    return num;
  }
  if (std::is_sorted(greatNumStore.rbegin(), greatNumStore.rend())) {
    return num;
  }
  std::sort(greatNumStore.begin(), greatNumStore.end(), std::greater<int>());
  int i = 1, j = greatNumStore.size() - 1;
  std::swap(greatNumStore[i], greatNumStore[j]);
  for (int digit : greatNumStore) {
    empty += std::to_string(digit);
  }
  return std::stoi(empty);
}
