#include <stdio.h>

int main(int argc, char *argv[]) {
  int i, j, m, n;
  int l = 9;
  printf("九九乘法表：\n\n");
  for (i = 1; i <= l; i++) {
    for (j = 1; j <= i; j++) { printf("%d*%d=%2d ", i, j, i * j); }
    printf("\n");
  }
  puts("");

  for (i = 1; i <= l; i++) {
    for (n = 1; n <= l - i; n++) { printf("    "); }
    for (j = 1; j <= i; j++) { printf("%d*%d=%2d  ", i, j, i * j); }
    printf("\n");
  }
  puts("");

  for (i = l; i > 0; i--) {
    for (j = i; j > 0; j--) { printf("%d*%d=%2d ", i, j, i * j); }
    printf("\n");
  }
  return 0;
}
