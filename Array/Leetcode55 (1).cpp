#include<vector>
bool canJump(std::vector<int>& nums)
{
  int goalpost = nums.size() - 1;
  for (int j = nums.size() - 1; j>=0; --j)
  {
    if (j + nums[j] >= goalpost)
    {
      goalpost = j;
    }
  }
    if (goalpost == 0)
    {
      return true;
    }
    else{
      return false;
    }

}
