#include <vector>
std::vector<int> searchRange(std::vector<int> &nums, int target) {
  std::vector<int> twoIndex(2, -1);
  if (nums.empty() || nums.size() <= 1) {
    return twoIndex;
  }
  for (int i = 0; i < nums.size(); ++i) {
    if (nums[i] == target) {
      twoIndex[0] = i;
      break;
    }
  }
  if (twoIndex[0] == -1) {
    return twoIndex;
  }
  for (int j = nums.size() - 1; j >= 0; --j) {
    if (nums[j] == target) {
      twoIndex[1] = j;
      break;
    }
  }
  return twoIndex;
}
