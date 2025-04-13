#include <string>

std::string predictPartyVictory(std::string senate) {
  int rCount{0};
  int dCount{0};
  std::string s = "Radiant";
  std::string str = "Dire";

  for (int i = 0; i < senate.length(); ++i) {
    if (senate[i] == 'R') {
      rCount++;
    }
    if (senate[i] == 'D') {
      dCount++;
    }
  }

  if (rCount > dCount) {
    return s;
  } else if (dCount == 0) {
    return s;
  } else if (rCount == 0) {
    return str;
  } else {
    return str;
  }
}
