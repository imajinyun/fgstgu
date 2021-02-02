#include <stdio.h>

int main(int argc, char *argv[]) {
  double rates[5] = {88.99, 77.12, 59.38, 11.29, 35.90};
  double *const ptr = rates; // ptr 指向数组的开头
  // ptr = &rates[2]; // 不允许，因为该指针不能指向别处

  // 可以用这种指针修改它所指向的值，但是它只能指向初始化时设置的地址
  *ptr = 99.99; // 没问题，更改了 rates[0] 的值
  printf("rates[0] = %f, *ptr = %f\n", rates[0], *ptr);
  for (int i = 0; i < 5; i++) {
    printf("rates[%d] = %f", i, *(ptr + i));
    printf("\n");
  }

  // 在创建指针时还可以使用 const 两次，该指针既不能更改它所指向的地址，也不能修改指向地址上的值
  const double *const p = rates;
  // p = &rates[2]     // 不允许
  // *p = 99.99; // 不允许
  return 0;
}
