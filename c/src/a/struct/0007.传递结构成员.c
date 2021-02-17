#include <stdio.h>

#define FUNDLEN 50

struct funds {
  char bank[FUNDLEN];
  double bankfund;
  char save[FUNDLEN];
  double savefund;
};

double sum(double, double);

/**
 * 传递结构成员。
 *
 * -> 只要结构成员是一个具有单个值的数据类型（即 int、char、float、double 或指针），
 *    便可把它作为参数传递给接受该特定类型的函数。
 */
int main(int argc, char *argv[]) {
  struct funds jack = {"ICBC", 4044.38, "Lucky's Savings and Loan", 8539.52};
  printf("Jack has a total of ￥%.2f\n", sum(jack.bankfund, jack.savefund));
  return 0;
}

double sum(double x, double y) { return (x + y); }
