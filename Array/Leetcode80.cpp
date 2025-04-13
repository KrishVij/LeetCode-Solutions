#include <vector>
int removeDuplicates(std::vector<int> &nums) {
  int ans_length = 0;
  int i = 0;
  int comparator = 1;
  int unique_counter = 0;
  for (int j = 0; j < nums.size(); ++j) {

    while (nums[i] != nums[comparator]) {
      nums[i] = nums[comparator];
      unique_counter++;
      i++;
      while (nums[i] == nums[comparator] && unique_counter <= 2) {
        nums[i] = nums[comparator];
        unique_counter++;
        i++;
      }
    }
    i++;
  }
  while (nums[i] == nums[comparator]) {
    ans_length++;
  }
  return ans_length;
}
