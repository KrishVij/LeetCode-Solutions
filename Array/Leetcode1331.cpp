#include <algorithm>
#include <vector>
std::vector<int> arrayRankTransform(std::vector<int> &arr) {
  std::vector<int> copy(arr);
  std::sort(copy.begin(), copy.end());
  std::vector<int> ans(arr.size());
  for (int i = 0; i < arr.size(); ++i) {
    for (int j = 0; j < copy.size(); ++j) {
      if (arr[i] == copy[j]) {
        ans[i] = j + 1;
        break;
      }
    }
  }

  for (int i = 1; i < copy.size(); ++i) {
    if (copy[i] == copy[i - 1]) {
      for (int k = 0; k < arr.size(); ++k) {
        if (arr[k] == copy[i]) {
          arr[k] = arr[k] - 1;
        }
      }
    }
  }
  return ans;
}
