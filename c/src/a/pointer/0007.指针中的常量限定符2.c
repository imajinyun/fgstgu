#include <stdio.h>

int main(int argc, char *argv[]) {
  // 指针所指向的数据是只读的，也就是 p1、p2 本身的值可以修改（指向不同的数据），但它们指向的数据不能被修改
  const int *p1;
  int const *p2;

  // 指针是只读的，也就是 p3 本身的值不能被修改，指针指向的值可以修改
  int *const p3;

  // 指针本身和它指向的数据都是只读的，既不能修改指针本身，也不能修改指针指向的数据
  const int *const p4;
  int const *const p5;

  printf("const 限定符出现在不同的位置!\n");
  return 0;
}
