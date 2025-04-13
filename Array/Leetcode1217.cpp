#include <vector>
int minCostToMoveChips(std::vector<int> &position) {
  int n = position.size();
  int evenCount = 0;
  int oddCount = 0;

  for (int i = 0; i < n; ++i) {
    if (position[i] % 2 == 0) {
      evenCount++;
    } else {
      oddCount++;
    }
  }
  return std::min(evenCount, oddCount);
}
