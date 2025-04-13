#include <vector>
std::vector<int> plusOne(std::vector<int> &digits) {
  std::vector<int> carry(digits.size());
  for (int i = digits.size() - 1; i >= 0; --i) {
    if (digits[i] != 9) {
      digits[i] += 1;
      return digits;

    } else if (digits[i] == 9) {
      digits[i] = 0;
    }
  }
  digits.insert(digits.begin(), 1);
  return digits;
}
