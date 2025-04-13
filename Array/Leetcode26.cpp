#include <vector>
int removeDuplicates(std::vector<int> &nums) {
  int j = 0;
  int k = 1;
  for (int i = 1; i < nums.size() && j < nums.size(); ++i) {
    if (nums[i] != nums[j]) {
      nums[j + 1] = nums[i];
      j++;
      k++;
    }
  }
  return k;
}
