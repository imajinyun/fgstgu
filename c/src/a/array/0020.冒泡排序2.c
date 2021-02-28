#include <stdbool.h>
#include <stdio.h>

int main(int argc, char *argv[]) {
  int nums[] = {4, 5, 2, 0, 7, 1, 8, 3, 6, 9};
  int i, j, tmp;
  int len = 10;
  bool flag;

  printf("冒泡排序前的结果：\n");
  for (i = 0; i < len; i++) { printf("%d ", nums[i]); }

  // 优化算法：最多进行 len-1 轮比较
  for (i = 0; i < len; i++) {
    flag = true; // 标志：假设剩下的元素已经排序好了
    for (j = 0; j < len - 1 - i; j++) {
      if (nums[j] > nums[j + 1]) {
        tmp = nums[j];
        nums[j] = nums[j + 1];
        nums[j + 1] = tmp;
        flag = false; // 一旦需要交换数组元素，就说明剩下的元素没有排序好
      }
    }

    if (flag) { // 如果没有发生交换，说明剩下的元素已经排序好了，退出就好
      break;
    }
  }
  puts("\n冒泡排序后的结果：");
  for (i = 0; i < len; i++) { printf("%d ", nums[i]); }
  printf("\n");
  return 0;
}
