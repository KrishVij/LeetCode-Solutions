#include <algorithm>
#include <vector>
int arrayPairSum(std::vector<int> &nums) {
  std::sort(nums.begin(), nums.end());
  int sum{0};
  int first{0};
  int second{1};

  if (nums.empty()) {
    return -1;
  }

  while (first < nums.size() && second < nums.size()) {
    sum += std::min(nums[first], nums[second]);

    first += 2;
    second += 2;
  }
  return sum;
}
