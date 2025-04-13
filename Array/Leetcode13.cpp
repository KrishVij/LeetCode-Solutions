#include <string>
int romanToInt(std::string s) {
  const std::string Roman[]{"I",  "1", "V",   "5", "X",   "10", "L",
                            "50", "C", "100", "D", "500", "M",  "100"};
  int n = sizeof(Roman) / sizeof(Roman[0]);
  int IntegerSum{0};
  for (int i = 0; i < s.length(); ++i) {
    std::string current_char(1, s[i]);
    int current_Value = 0;
    for (int j = 0; j < n; j += 2) {
      if (current_char == Roman[j]) {
        current_Value = std::stoi(Roman[j + 1]);
        break;
      }
    }

    int next_value = 0;
    if (i + 1 < s.length()) {
      std::string next_char(1, s[i + 1]);
      for (int k = 0; k < n; k += 2) {
        if (next_char == Roman[k]) {
          next_value = std::stoi(Roman[k + 1]);
          break;
        }
      }
    }
    if (next_value > current_Value) {
      IntegerSum -= current_Value;
    } else {
      IntegerSum += current_Value;
    }
  }
  return IntegerSum;
}
