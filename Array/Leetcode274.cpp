#include <algorithm>
#include <vector>
int hIndex(std::vector<int> &citations) {
  std::sort(citations.begin(), citations.end());
  for (int i = 0; i < citations.size(); ++i) {
    if (citations.size() - i <= citations[i]) {
      return citations.size() - i;
    }
  }
  return -1;
}
