#include <string>
#include <vector>
bool reportSpam(std::vector<std::string> &message,
                std::vector<std::string> &bannedWords) {
  int count{0};
  if (message.empty() || bannedWords.empty()) {
    return false;
  }
  for (int i = 0; i < message.size(); ++i) {
    for (int j = 0; j < bannedWords.size(); ++j) {
      if (bannedWords[j] == message[i]) {
        ++count;
        break;
      }
    }
    if (count >= 2) {
      return true;
    }
  }
  return false;
}
