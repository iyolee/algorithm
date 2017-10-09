/**
问题描述：
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

例子：
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;
class Solution {
  public:
    vector<int> twoSum(vector<int>& nums, int target) {
      vector<int> num = nums;
      std::sort(num.begin(), num.end());

      int length = nums.size();
      int left = 0;
      int right = length - 1;
      int sum = 0;

      vector<int> index;
      while (left < right) {
        sum = num[left] + num[right];

        if (sum == target) {
          for (int i=0; i<length; i++) {
            if (nums[i] == num[left]) {
              index.push_back(i);
            } else if (nums[i] == num[right]) {
              index.push_back(i);
              }
              if (index.size() == 2) {
                break;
              }
          }
            break;
        } else if (sum > target) {
          --right;
          } else if (sum < target) {
            ++left;
          }
      }
      return index;
    }
};

int main() {
  Solution p;
  int nums[] = {2, 7, 11, 15};
  vector<int> a(nums, nums + 4);
  int t = 9;
  vector<int> r;
  r = p.twoSum(a, t);
  for (int i = 0; i<r.size(); i++) {
    cout << "index" << i+1 << "=" <<r[i] <<endl;
  }
  return 0;
}