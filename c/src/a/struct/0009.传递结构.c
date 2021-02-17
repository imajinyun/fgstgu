#include <stdio.h>

#define FUNDLEN 50

struct funds {
  char bank[FUNDLEN];
  double bankfund;
  char save[FUNDLEN];
  double savefund;
};

double sum(struct funds fs);

int main(int argc, char *argv[]) {
  struct funds jack = {"ICBC", 4096.32, "Lucky's Savings and Loan", 8192.64};
  printf("Jack has a total of ￥%.2f\n", sum(jack));
  return 0;
}

double sum(struct funds fs) { return (fs.bankfund + fs.savefund); }
