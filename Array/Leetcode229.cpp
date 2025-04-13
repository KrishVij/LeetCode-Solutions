#include <algorithm>
#include <vector>
std::vector<int> majorityElement(std::vector<int> &nums) {
  std::vector<int> singleLe;
  std::vector<std::pair<int, int>> hiFreq;
  std::sort(nums.begin(), nums.end());
  int n = nums.size() / 3;
  int count{1};
  int index{0};
  if (nums.empty()) {
    return {-1};
  }
  while (index < nums.size() - 1) {
    if (nums[index] == nums[index + 1]) {
      count++;
    } else {
      hiFreq.emplace_back(nums[index],
                          count); // Store the element and its count
      count = 1;
    }
    index++;
  }
  hiFreq.emplace_back(nums[index], count);
  for (const auto &ele : hiFreq) {
    if (ele.second > n) {
      singleLe.emplace_back(ele.first);
    }
  }
  return singleLe.empty() ? std::vector<int>{} : singleLe;
}
