#include <stdio.h>

#define LENGTH 21

struct names {
  char first[LENGTH];
  char last[LENGTH];
};

struct guy {
  struct names name;
  char favfood[LENGTH];
  char job[LENGTH];
  float income;
};

int main(int argc, char *argv[]) {
  struct guy fellow[2] = {
    {{"Even", "Villard"}, "grilled salmon", "personality coach", 68112.00},
    {{"Rodney", "Swillbelly"}, "tripe", "tabloid editor", 433200.00},
  };
  struct guy *him; // 指向结构的指针

  printf("address #1: %p, #2: %p\n", &fellow[0], &fellow[1]);
  him = &fellow[0];
  printf("pointer #1: %p #2, %p\n", him, him + 1);
  printf("him->income is $%.2f, (*him).income is $%.2f\n", him->income, (*him).income);
  him++;
  printf("him->favfood is %s, him->name.last is %s\n", him->favfood, him->name.last);
  return 0;
}
