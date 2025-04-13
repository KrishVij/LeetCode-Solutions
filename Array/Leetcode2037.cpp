#include <algorithm>
#include <cstdlib>
#include <vector>
int minMovesToSeat(std::vector<int> &seats, std::vector<int> &students) {
  std::sort(seats.begin(), seats.end());
  std::sort(students.begin(), students.end());
  std::vector<int> difference;
  int sum = 0;

  for (int i = 0, j = 0; i < seats.size() && j < students.size(); ++i, ++j) {
    difference.emplace_back(std::abs(seats[i] - students[j]));
  }
  for (int k = 0; k < difference.size(); ++k) {
    sum += difference[k];
  }
  return sum;
}
