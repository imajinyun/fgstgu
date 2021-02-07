#include <stdio.h>

int main(int argc, char *argv[]) {
  int i = 1;
  int j;

  printf("跳出循环：\n");
  while (1) {
    j = 1;
    while (1) {
      printf("%4d", i * j);
      j++;
      if (j >= 5) { break; }
    }
    printf("\n");
    i++;

    if (i >= 5) { break; }
  }
  return 0;
}
