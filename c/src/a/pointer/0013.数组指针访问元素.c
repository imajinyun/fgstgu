#include <stdio.h>

int main(int argc, char *argv[]) {
  int arr[] = {3, 2, 1, 5, 4};
  int len = sizeof(arr) / sizeof(int); // 求数组长度
  int i;
  int *p = &arr[2]; // 也可以写成：int *p = arr + 2;

  printf("指针访问数组中的元素：\n");
  printf("%d %d %d %d %d\n", *(p - 2), *(p - 1), *p, *(p + 1), *(p + 2));

  int *q = &arr;
  printf("\n自增运算符来遍历数组元素：\n");
  for (i = 0; i < len; i++) { printf("%d ", *q++); }
  puts("\n\n注意：假设 p 是指向数组 arr 中第 n 个元素的指针，那么 *p++、*++p、(*p)++ 分别是什么意思：\n"
       "1. *p++：等价于 *(p++)，表示先取得第 n 个元素的值，再将 p 指向下一个元素；\n"
       "2. *++p：等价于 *(++p)，会先进行 ++p 运算，使得 p 的值增加，指向下一个元素，"
       "整体上相当于 *(p+1)，所以会获得第 n+1 个数组元素的值；\n"
       "3. (*p)++：先取得第 n 个元素的值，再对该元素的值加 1。假设 p 指向第 0 个元素，并且第 0 个元素的值为 99，"
       "执行完该语句后，第 0 个元素的值就会变为 100；");
  return 0;
}
