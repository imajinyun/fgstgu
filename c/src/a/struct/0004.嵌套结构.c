#include <stdio.h>

#define LENGTH 21

const char *msgs[5] = {"Thank you for the wonderful evening, ", "You certainly prove that a ",
                       "is a special kind of guy. We must get together ", "over a delicious ",
                       " and have a few laughs"};

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
  struct guy fellow = {{"Ewen", "Villard"}, "grilled salmon", "personality coach", 76112.00};
  printf("Dear %s, \n\n", fellow.name.first);
  printf("%s%s\n", msgs[0], fellow.name.first);
  printf("%s%s\n", msgs[1], fellow.job);
  printf("%s\n", msgs[2]);
  printf("%s%s%s", msgs[3], fellow.favfood, msgs[4]);
  if (fellow.income > 150000.0) {
    puts("!!!");
  } else if (fellow.income > 75000.00) {
    puts("!!");
  } else {
    puts("!");
  }
  printf("\n%40s%s\n", " ", "See you soon,");
  printf("%40s%s\n", " ", "Shalala");
  return 0;
}
