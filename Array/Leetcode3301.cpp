#include <algorithm>
#include <vector>
long long maximumTotalSum(std::vector<int> &maximumHeight) {

  std::sort(maximumHeight.begin(), maximumHeight.end());
  std::vector<int> copyHeight(maximumHeight.size());
  long long totalSum{1};
  if (maximumHeight.empty() || maximumHeight.size() <= 2) {
    return -1;
  }
  for (int i = 0; i < maximumHeight.size(); ++i) {
    if (maximumHeight[i] == maximumHeight[i + 1]) {
      return -1;
    }
  }
  copyHeight = maximumHeight;

  if (copyHeight[0] > 1) {
    copyHeight[0]--;
  } else {
    return -1;
  }
  for (int i = 0; i < copyHeight.size(); ++i) {
    if (i == 0) {
      --copyHeight[i];
    }
    totalSum += copyHeight[i];
  }
  return totalSum;
}
