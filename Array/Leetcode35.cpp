#include <vector>
int searchInsert(std::vector<int> &nums, int target) {
  int pos = 1;
  for (int i = 0; i < nums.size(); ++i) {
    if (nums[i] == target) {
      return i;
    } else if (pos == target) {
      return pos - 1;
    } else {
      pos++;
    }
  }
}
