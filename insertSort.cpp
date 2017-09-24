//插入排序
#include <iostream>
int main() {
  int arr[] = {66, 6, 666, -6, 66666, 6666};
  for (int i = 0; i < sizeof(arr) / sizeof(arr[0]); i++) {
    int index = 0;
    for (int j = 0; j < sizeof(arr) / sizeof(arr[0]) - i; j++) {
      if (arr[j] > arr[index]) {
        index = j;
      }
    }
    int temp = arr[sizeof(arr) / sizeof(arr[0]) - i - 1];
    arr[sizeof(arr) / sizeof(arr[0]) - i - 1] = arr[index];
    arr[index] = temp;
  }
  for (int t = 0; t < sizeof(arr) / sizeof(arr[0]); t++) {
    std::cout << arr[t] << " ";
  }
  return 0;
}
