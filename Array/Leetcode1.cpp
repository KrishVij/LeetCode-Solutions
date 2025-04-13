#include<vector>
std::vector<int> twoSum(std::vector<int>& nums, int target)
{
  std::vector<int> twoNums;
  if (nums.size() <= 1)
  {
    return nums;
  }
  for (int i = 0;i<nums.size();++i)
  {
    for (int j = i+1;j<nums.size();++j)
    {
      if (nums[i] + nums[j] == target)
      {
          twoNums.push_back(i);
          twoNums.push_back(j);
      
      }
    }
   
  }
  return twoNums;

}
