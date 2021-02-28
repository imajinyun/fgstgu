#include <stdio.h>

typedef char (*PTR_TO_ARR)[30];
typedef int (*PTR_TO_FUNC)(int, int);

int max(int a, int b) { return a > b ? a : b; }

char str[3][30] = {
  "https://github.com/",
  "C Language",
  "C++ Language",
};

int main(int argc, char *argv[]) {
  PTR_TO_ARR parr = str;
  PTR_TO_FUNC pfn = max;
  int i;
  printf("Max :%d\n", (*pfn)(88, 99));
  for (i = 0; i < 3; i++) { printf("str[%d]: %s\n", i, *(parr + i)); }
  return 0;
}
