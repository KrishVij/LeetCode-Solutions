#include <algorithm>
#include <vector>
using namespace std;
void merge(vector<int> &nums1, int m, vector<int> &nums2, int n) {
  int index = 0;
  if (n == 0) {
  }
  if (m == 0) {
    for (int i = 0; i < nums1.size(); i++) {
      nums1[i] = nums2[i];
    }
  }
  int nums2_length = nums2.size();
  for (int i = m; i < nums1.size() && index < n; ++i) {
    nums1[i] = nums2[index];
    index++;
  }
  sort(nums1.begin(), nums1.end());
}
