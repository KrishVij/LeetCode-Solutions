#include <algorithm>
#include <cctype>
#include <regex>
#include <string>
bool isPalindrome(std::string s) {
  if (s == " ") {
    return true;
  }
  if (s.empty()) {
    return false;
  }
  std::regex non_alnum_regex("[^a-zA-Z0-9]");
  std::string onlyWords = std::regex_replace(s, non_alnum_regex, "");
  std::transform(onlyWords.begin(), onlyWords.end(), onlyWords.begin(),
                 ::tolower);
  std::string pallindromeTest = onlyWords;
  std::transform(onlyWords.begin(), onlyWords.end(), onlyWords.begin(),
                 ::tolower);
  std::reverse(pallindromeTest.begin(), pallindromeTest.end());
  if (pallindromeTest == onlyWords) {
    return true;
  } else {
    return false;
  }
}
