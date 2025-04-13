#include <vector>
int canCompleteCircuit(std::vector<int> &gas, std::vector<int> &cost) {
  int start = 0;
  int g = gas.size();
  int c = cost.size();
  int gasLeft = 0;
  int totalgas = 0;
  for (int i = 0; i < g; ++i) {
    int diff  = gas[i] - cost[i];
    gasLeft += diff;
    totalgas += diff;
    if (gasLeft < 0) {
      start = i+1;
      gasLeft = 0;
    }
    
  }
  if (totalgas < 0)
  {
    return -1;
  }
  return start;
}
