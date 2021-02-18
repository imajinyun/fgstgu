#include <stdio.h>

union hold {
  int digit;
  double bigger;
  char letter;
};

void test1() {
  printf("union initialized!\n");
  union hold fit;      // hold 类型的联合变量
  union hold save[10]; // 内含 10 个联合变量的数组
  union hold *pu;      // 指向 hold 类型联合变量的指针

  union hold a;
  a.letter = 'A';
  union hold b = a;                 // 用另一个联合来初始化
  union hold c = {88};              // 初始化联合的 digit 成员
  union hold d = {.bigger = 99.99}; // 指定初始化器
}

void test2() {
  puts("1. 点运算符表示正在使用哪种数据类型。在联合中，一次只存储一个值。\n"
       "2. 即使有足够的空间，也不能同时存储一个 char 类型值和一个 int 类型值。\n"
       "3. 编写代码时要注意当前存储在联合中的数据类型。");
  union hold uh;
  uh.digit = 88;   // 把 88 存储在 hu，占 2 个字节
  uh.bigger = 4.2; // 清除 88，存储 4.2，占 8 个字节
  uh.letter = 'H'; // 清除 4.2，存储 H，占 1 个字节

  union hold *pu;
  pu = &uh;
  int x = pu->digit; // 相当于 int x = uh.digit;
}

int main(int argc, char *argv[]) {
  test1();
  printf("\n");
  test2();
  return 0;
}
