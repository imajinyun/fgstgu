#include <stdio.h>

int main(int argc, char *argv[]) {
  enum weeks { Mon = 1, Tue, Wed, Thu, Fri, Sat, Sun } day = Mon;
  printf("%lu %lu %lu %lu %lu\n", sizeof(enum weeks), sizeof(day), sizeof(Mon), sizeof(Wed), sizeof(int));
  return 0;
}
