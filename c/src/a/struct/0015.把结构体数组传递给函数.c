#include <stdio.h>

#define FUNDLEN 50
#define N 2

struct funds {
  char bank[FUNDLEN];
  double bankfund;
  char save[FUNDLEN];
  double savefund;
};

double sum(const struct funds fs[], int n);

int main(int argc, char *argv[]) {
  struct funds fs[N] = {
    {"Garlic-Melon Bank", 4023.27, "Lucky's Savings and Loan", 8848.89},
    {"Honest Jack's Bank", 3629.88, "Party Time Savings", 2989.91},
  };
  printf("The Jhonses have a total of ￥%.2f\n", sum(fs, N));
  return 0;
}

double sum(const struct funds fs[], int n) {
  double total;
  int i;
  for (i = 0, total = 0; i < n; i++) { total += fs[i].bankfund + fs[i].savefund; }
  return total;
}
