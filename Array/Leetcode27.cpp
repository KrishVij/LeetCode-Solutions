#include <vector>
void swap(int *first_pointer, int *second_pointer) {
  int temp = *first_pointer;
  *first_pointer = *second_pointer;
  *second_pointer = temp;
}

int removeElement(std::vector<int> &nums, int val) {
  int swap_counter = 0;
  int index = 0;
  int first_element = 0;
  for (int i = 0; i < nums.size(); ++i) {
    if (nums[i] == val) {
      swap(&nums[index], &nums[i]);
      index++;
      swap_counter++;
    }
  }
  return swap_counter;
}
