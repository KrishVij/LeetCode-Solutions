#include <algorithm>
#include <vector>
std::vector<std::vector<int>> threeSum(std::vector<int> &nums) {
  std::vector<std::vector<int>> triplet;
  std::sort(nums.begin(), nums.end());
  if (nums.empty()) {
    return {};
  }

  for (int i = 0; i < nums.size(); ++i) {
    for (int j = i + 1; j < nums.size(); ++j) {
      for (int k = j + 1; k < nums.size(); ++k) {
        if (i != j && j != k && j != k && nums[i] + nums[j] + nums[k] == 0) {
          triplet.push_back({nums[i], nums[j], nums[k]});
        }
      }
    }
  }
  return triplet;
}
