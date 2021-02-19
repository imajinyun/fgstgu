#include <stdio.h>

int main(int argc, char *argv[]) {
  int i = 100;

  /**
   * 1. p 保存了 i 的地址，所以 p 指向了 i；
   * 2. p 不是 i，i 也不是 p，也就是说，修改 p 的值不影响 i 的值，反之亦然；
   * 3. 一个指针变量指向了某个普通变量，则 *指针变量不等价于普通变量；
   * 4. *p 就是以 p 的内容为地址的变量；
   */
  int *p = &i;
  printf("i = %d, *p = %d\n", i, *p);
  puts("注意：使用指针是间接获取数据，使用变量名是直接获取数据，前者比后者的代价要高");
  return 0;
}
