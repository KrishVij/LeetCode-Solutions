#include <vector>
bool canArrange(std::vector<int> &arr, int k) {
  bool isDivisible = false;
  int count = 0;
  for (int i = 0; i < arr.size(); ++i) {
    for (int j = i + 1; j < arr.size(); ++j) {
      if ((arr[i] + arr[j] % k) == 0) {
        isDivisible = true;
        count++;
      } else {
        isDivisible = false;
      }
    }
  }
  if (count == arr.size() / 2) {
    return true;
  }
  return false;
}
