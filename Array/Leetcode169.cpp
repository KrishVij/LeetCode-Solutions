#include <algorithm>
#include <vector>
int majorityElement(std::vector<int> &nums) {
  std::vector<int> countStorage;
  countStorage.resize(2 * nums.size());
  int ans = 0;
  int i = 0;
  int comparator = 1;
  int index = 0;
  int majority_count = 1;
  std::sort(nums.begin(), nums.end());
  while (comparator < nums.size()) {
    if (nums[i] == nums[comparator]) {
      majority_count++;

    } else {
      countStorage.push_back(nums[i]);
      countStorage.push_back(majority_count);
      i = comparator;
      majority_count = 1;
    }
    comparator++;
  }
  countStorage.push_back(nums[i]);
  countStorage.push_back(majority_count);
  for (int j = 1; j < countStorage.size(); j += 2) {
    if (countStorage[j] > nums.size() / 2) {
      ans = countStorage[j - 1];
    }
  }
  return ans;
}
