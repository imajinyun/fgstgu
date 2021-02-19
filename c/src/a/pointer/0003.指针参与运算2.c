#include <stdio.h>

int main(int argc, char *argv[]) {
  int a = 10, *pa = &a, *paa = &a;
  double b = 9.9, *pb = &b;
  char c = 'E', *pc = &c;

  printf("最初的地址值：\n");
  printf("&a = %#X, &b = %#X, &c = %#X\n", &a, &b, &c);
  printf("pa = %#X, pb = %#X, pc = %#X\n", pa, pb, pc);
  printf("paa = %#X, *paa = %d\n", paa, *paa);

  printf("\n加法运算：\n");
  pa++;
  pb++;
  pc++;
  printf("&a = %#X, &b = %#X, &c = %#X\n", &a, &b, &c);
  printf("pa = %#X, pb = %#X, pc = %#X\n", pa, pb, pc);
  printf("paa = %#X, *paa = %d\n", paa, *paa);

  printf("\n减法运算：\n");
  pa -= 2;
  pb -= 2;
  pc -= 2;
  printf("&a = %#X, &b = %#X, &c = %#X\n", &a, &b, &c);
  printf("pa = %#X, pb = %#X, pc = %#X\n", pa, pb, pc);
  printf("paa = %#X, *paa = %d\n", paa, *paa);

  printf("\n比较运算：\n");
  printf("pa = %#X, paa = %#X\n", pa, paa);

  if (pa == paa) {
    printf("*paa = %d\n", *paa);
  } else {
    printf("*pa = %d\n", *pa);
  }

  printf("\n");
  puts("注意：\n"
       "1. 不要对指向普通变量的指针进行加减运算；\n"
       "2. 不能对指针变量进行乘法、除法、取余等其他运算，除了会发生语法错误，也没有实际的含义；");
  return 0;
}
