#include <stdio.h>

#define FUNDLEN 50

struct funds {
  char bank[FUNDLEN];
  double bankfund;
  char save[FUNDLEN];
  double savefund;
};

double sum(const struct funds *); // 参数是一个指针

int main(int argc, char *argv[]) {
  struct funds jack = {"ICBC", 4055.72, "Lucky's Savings and Loan", 8239.32};
  printf("Jack has a total of ￥%.2f\n", sum(&jack));
  return 0;
}

double sum(const struct funds *fs) { return (fs->bankfund + fs->savefund); }
