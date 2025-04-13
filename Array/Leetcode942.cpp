#include <string>
#include <vector>
std::vector<int> diStringMatch(std::string s) {
  int n = s.length();
  std::vector<int> perm(n + 1);
  std::vector<int> ans(n + 1);

  if (s.empty()) {
    return {};
  }

  for (int i = 0; i <= n; ++i) {
    perm[i] = i;
  }

  int left{0};
  int right = perm.size() - 1;

  for (int j = 0, k = 0; j < perm.size() && k < n; ++j, ++k) {
    if (left < perm.size() && right >= 0) {
      if (s[k] == 'I') {
        ans[k] = perm[left++];
      } else if (s[k] == 'D') {
        ans[k] = perm[right--];
      }
    }
  }
  ans[n] = perm[left];
  return ans;
}
