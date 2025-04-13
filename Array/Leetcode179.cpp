#include <algorithm>
#include <string>
#include <vector>
std::string largestNumber(std::vector<int> &nums) {
  std::sort(nums.begin(), nums.end());
  std::string largestNum = "";

  if (nums.empty()) {
    return largestNum;
  }

  for (int i = 0; i < nums.size(); ++i) {
    largestNum += std::to_string(nums[i]);
  }
  return largestNum;
}
