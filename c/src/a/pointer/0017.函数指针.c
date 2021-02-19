#include <stdio.h>

int *func();

int main(int argc, char *argv[]) {
  int *p = func(), n;
  n = *p;

  puts("注意：\n"
       "func() 运行结束后 n 的内存依然保持原样，值还是 1234，如果使用及时也能够得到正确的数据，\n"
       "如果有其它函数被调用就会覆盖这块内存，得到的数据就失去了意义。");
  printf("\nn = %d\n", n);
  return 0;
}

int *func() {
  int n = 1234;

  // Note: address of stack memory associated with local variable 'n' returned.
  return &n;
}
