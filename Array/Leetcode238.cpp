#include <vector>
std::vector<int> productExceptSelf(std::vector<int> &nums) {
  std::vector<int> productVector(nums.size(), 1);
  int leftproduct = 1;
  for (int i = 0; i < nums.size(); ++i) {
    productVector[i] = leftproduct;
    leftproduct *= nums[i];
  }
  int rigthproduct = 1;
  for (int i = nums.size() - 1; i >= 0; --i) {
    productVector[i] *= rigthproduct;
    rigthproduct *= nums[i];
  }
  return productVector;
}
