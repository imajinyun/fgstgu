#include <stdio.h>

int a[2] = {100, 200};
int b[2] = {300, 400};

// 一元运算符 * 和 ++ 的优先级相同，但结合律是从右往左
int main(int argc, char *argv[]) {
  int *p1, *p2, *p3;
  p1 = p2 = a;
  p3 = b;
  printf("*p1 = %d, *p2 = %d, *p3 = %d\n", *p1, *p2, *p3);

  // *p1++ 和 *p2++ 分别把 p1 和 p2 指向数组的下一个元素，(*p3)++ 改变了数组的元素的值
  printf("*p1++ = %d, *p2++ = %d, (*p3)++ = %d\n", *p1++, *p2++, (*p3)++);
  printf("*p1 = %d, *p2 = %d, *p3 = %d\n", *p1, *p2, *p3);
  return 0;
}
