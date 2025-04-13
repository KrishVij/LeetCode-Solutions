#include <string>
int minSwaps(std::string s) {
  int balance{0};
  int swap{0};
  if (s.empty()) {
    return 0;
  }
  for (auto ch : s) {
    if (ch == ']') {

      --balance;

    } else if (ch == '[') {

      ++balance;
    }
    if (balance < 0) {
      ++swap;
      balance = 1;
    }
  }

  return swap;
}
