#include <cstddef>
#include <string>
#include <vector>
std::string longestCommonPrefix(const std::vector<std::string> &strs) {
  std::string prefix = strs[0];
  if (strs.empty()) {
    return "";
  }
  for (size_t i = 1; i < strs.size(); ++i) {
    while (strs[i].find(prefix) != 0) {
      prefix = prefix.substr(0, prefix.size() - 1);
      if (prefix.empty()) {
        return "";
      }
    }
  }
  return prefix;
}
