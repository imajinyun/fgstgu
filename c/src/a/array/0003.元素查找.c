#include <stdio.h>

#define LEN 15

int main(int argc, char *argv[]) {
  int num;
  int nums[LEN] = {0, 1, 7, 8, 18, 29, 38, 40, 45, 60, 77, 80, 84, 97, 99};
  int index = -1;

  printf("请输入一个数字：");
  scanf("%d", &num);
  for (int i = 0; i < LEN; i++) {
    if (nums[i] == num) {
      index = i;
      break;
    } else if (nums[i] > num) {
      // nums[i] 后边的元素也都大于 num 了，num 肯定不在数组中了，就没有必要再继续比较了
      break;
    }
  }

  if (index > 0) {
    printf("%d 元素在数组中，索引为 %d\n", num, index);
  } else {
    printf("%d 元素不在数组中\n", num);
  }
  return 0;
}
