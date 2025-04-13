#include <vector>
std::vector<int> getSneakyNumbers(std::vector<int> &nums) {
  std::vector<int> twoNumbers;
  int index = 0;
  if (nums.empty()) {
    return nums;
  }

  for (int i = 0; i < nums.size(); ++i) {
    for (int j = 0; j < nums.size(); ++j) {
      if (nums[i] == nums[j]) {
        if (index < 2) {
          twoNumbers[index] = nums[i];
          ++index;
        }
      }
    }
  }
  return twoNumbers;
}
