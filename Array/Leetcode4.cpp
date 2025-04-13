#include <algorithm>
#include <vector>
double findMedianSortedArrays(std::vector<int> &nums1,
                              std::vector<int> &nums2) {
  std::vector<int> mergedArray(nums1.size() + nums2.size());
  int m = mergedArray.size();
  for (int i = 0; i < nums1.size(); ++i) {
    mergedArray[i] = nums1[i];
  }
  for (int j = 0; j < nums2.size(); ++j) {
    mergedArray[nums1.size() + j] = nums2[j];
  }
  std::sort(nums1.begin(), nums2.end());
  if (m % 2 != 0) {
    int median = mergedArray[(m + 1) / 2];
    return median;
  } else {
    double median = (mergedArray[m / 2] + mergedArray[m / 2 - 1]) / 2.0;
    return median;
  }
}
