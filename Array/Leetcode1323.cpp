#include <algorithm>
#include <vector>
int maximum69Number(int num) {
  std::vector<int> digits;
  int sum = 0;

  while (num != 0) {
    int x = num % 10;
    digits.emplace_back(x);
    num = num / 10;
  }
  std::reverse(digits.begin(), digits.end());

  for (int i = 0; i < digits.size(); ++i) {
    if (digits[i] == 6) {
      digits[i] = 9;
    }
  }
  for (int digit : digits) {
    sum = sum * 10 + digit;
  }
  return sum;
}
