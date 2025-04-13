#include <string>
long long minimumSteps(std::string s) {
  int oneCounter{0};
  int zeroCounter{0};
  int r = s.length();

  for (int j = 0; j < r; ++j) {
    if (s[j] == '1') {
      oneCounter++;
    }
  }
  else if (s[j] == '0') {
    zeroCounter += oneCounter;
  }
  return zeroCounter;
}
