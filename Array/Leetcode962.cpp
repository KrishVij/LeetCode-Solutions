#include <vector>
int maxWidthRamp(std::vector<int> &nums) {
  int rampWidth{0};
  int maximum{0};
  int start{0};
  int end{0};

  if (nums.empty()) {
    return -0;
  }

  for (int i = 0; i < nums.size(); ++i) {
    for (int j = i + 1; j < nums.size(); ++j) {
      if (i < j && nums[i] <= nums[j]) {
        rampWidth = j - i;
        if (rampWidth > maximum) {
          maximum = rampWidth;
        }
      }
    }
  }
  return maximum;
}
