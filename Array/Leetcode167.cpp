#include <vector>
std::vector<int> twoSum(std::vector<int> &numbers, int target) {
  if (numbers.size() <= 1) {
    return {-1, -1};
  }
  for (int i = 0; i < numbers.size(); ++i) {
    for (int j = i + 1; j < numbers.size(); ++j) {
      if (numbers[i] + numbers[j] == target) {
        return {i + 1, j + 1};
      }
    }
  }
  return {-1, -1};
}
