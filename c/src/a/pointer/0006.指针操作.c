#include <stdio.h>

/**
 * 指针操作。
 *
 * 解引用：* 运算符给出指针指向地址上存储的值。
 * 取地址：和所有变量一样，指针变量也有自己的地址和值。对指针而言，& 运算符给出指针本身的地址。
 * 指针与整数相加：可以使用 + 运算符把指针与整数相加，或整数与指针相加。
 *   无论哪种情况，整数都会和指针所指向类型的大小（以字节为单位）相乘，然后把结果与初始地址相加。
 * 递增指针：递增指向数组元素的指针可以让该指针移动至数组的下一个元素；
 * 指针与整数相减：可以使用 - 运算符从一个指针中减去一个整数。
 *   指针必须是第 1 个运算对象，整数是第 2 个运算对象。
 *   该整数将乘以指针指向类型的大小（以字节为单位），然后用初始地址减去乘积。
 * 递减指针：递减指向数组元素的指针可以让该指针移动至数组的上一个元素；
 * 指针求差：可以计算两个指针的差值。通常，求差的两个指针分别指向同一个数组的不同元素，
 *   通过计算求出两元素之间的距离。差值的单位与数组类型的单位相同。
 * 指针比较：使用关系运算符可以比较两个指针的值，前提是两个指针都指向相同类型的对象。
 *
 * @param argc
 * @param argv
 * @return
 */
int main(int argc, char *argv[]) {
  int arr[5] = {100, 200, 300, 400, 500};
  int *p1, *p2, *p3;

  p1 = arr; // 把一个地址赋给指针
  p2 = &arr[2]; // 把一个地址赋给指针解引用指针，以及获取指针的地址
  printf("pointer value, dereferenced pointer, pointer address:\n");
  printf("p1 = %p, *p1 = %d, &p1 = %p\n", p1, *p1, &p1);

  p3 = p1 + 4;
  printf("\nadding an int to a pointer:\n");
  printf("p1 + 4 = %p, *(p1 + 4) = %d\n", p1 + 4, *(p1 + 4));

  p1++;
  printf("\nvalues after p1++:\n");
  printf("p1 = %p, *p1 = %d, &p1 = %p\n", p1, *p1, &p1);

  p2--;
  printf("\nvalues after p2--:\n");
  printf("p2 = %p, *p2 = %d, &p2 = %p\n", p2, *p2, &p2);

  // 恢复为初始值
  --p1;
  ++p2;

  printf("\nPointers reset to original values:\n");
  printf("p1 = %p, p2 = %p\n", p1, p2);

  // 一个指针减去另一个指针
  printf("\nSubtracting one pointer from another:\n");
  printf("p1 = %p, p2 = %p, p2 - p1 = %td\n", p1, p2, p2 - p1);
  printf("p1 = %d, p2 = %d, *p2 - *p1 = %d\n", *p1, *p2, *p2 - *p1);

  // 一个指针减去一个整数
  printf("\nSubtracting an int from a pointer:\n");
  printf("p3 = %p, p3 - 2 = %p, p3 = %d, *(p3 - 2) = %d\n", p3, p3 - 2, *p3, *(p3 - 2));
  return 0;
}
