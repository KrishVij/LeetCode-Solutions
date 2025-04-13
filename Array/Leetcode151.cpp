#include <string>
#include <vector>
std::string reverseWords(std::string s) {
  std::vector<char> stringToChar;
  int sL = s.length();
  std::string newStr;
  int start = 0;
  int end = stringToChar.size() - 1;
  if (s.length() <= 1) {
    return s;
  }
  while (start < sL && s[start] == ' ') {
    ++start;
  }
  while (end >= 0 && s[end] == ' ') {
    --end;
  }
  for (int i = start; i <= end; ++i) {
    if (s[i] != ' ' || s[i] == ' ' && s[i + 1] != ' ') {
      stringToChar.push_back(s[i]);
    }
  }
  while (start < end) {
    int temp = stringToChar[start];
    stringToChar[start] = stringToChar[end];
    stringToChar[end] = temp;
    ++start;
    --end;
  }
  start = 0;
  for (int i = 0; i <= stringToChar.size(); ++i) {

    if (i == stringToChar.size() || stringToChar[i] == ' ') {
      end = i - 1;
      while (start < end) {
        int temp = stringToChar[start];
        stringToChar[start] = stringToChar[end];
        stringToChar[end] = temp;
        ++start;
        --end;
      }
      start = i + 1;
    }
  }
  for (int i = 0; i < stringToChar.size(); ++i) {
    if (!(i > 0 && stringToChar[i] == ' ' && stringToChar[i - 1] == ' ')) {
      newStr.push_back(stringToChar[i]);
    }
  }

  if (!newStr.empty() && newStr.back() == ' ') {
    newStr.pop_back();
  }

  return newStr;
}
