#include <stdio.h>

int main(int argc, char *argv[]) {
  int arr[] = {3, 2, 1, 5, 4};
  int len = sizeof(arr) / sizeof(int);
  int i;

  printf("数组指针来遍历数组元素：\n");
  for (i = 0; i < len; i++) {
    // *(arr + i) 等价于 arr[i]
    printf("%d ", *(arr + i));
  }
  printf("\n");

  int *p = arr;
  printf("指针的方式遍历数组元素：\n");
  for (i = 0; i < len; i++) { printf("%d ", *(p + i)); }
  printf("\n");
  return 0;
}
