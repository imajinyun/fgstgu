#include <stdio.h>

int main(int argc, char *argv[]) {
  int nums[] = {4, 5, 2, 0, 7, 1, 8, 3, 6, 9};
  int i, j, tmp;
  int len = 10;

  printf("冒泡排序前的结果：\n");
  for (i = 0; i < len; i++) { printf("%d ", nums[i]); }

  for (i = 0; i < len; i++) {
    for (j = 0; j < len - 1 - i; j++) {
      if (nums[j] > nums[j + 1]) {
        tmp = nums[j];
        nums[j] = nums[j + 1];
        nums[j + 1] = tmp;
      }
    }
  }
  puts("\n冒泡排序后的结果：");
  for (i = 0; i < len; i++) { printf("%d ", nums[i]); }
  printf("\n");
  return 0;
}
