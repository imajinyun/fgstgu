#include <stdio.h>

void output1(const char *);
int output2(const char *);

int main(int argc, char *argv[]) {
  output1("If I'd as mush money");
  output1(" as I could spend,\n");
  printf("I count %d characters.\n", output2("I never would cry old chairs to mend"));
  return 0;
}

void output1(const char *str) {
  // 与 *str != '\0' 相同
  while (*str) { putchar(*str++); }
}

int output2(const char *str) {
  int count = 0;
  while (*str) {
    putchar(*str++);
    count++;
  }
  putchar('\n');
  return count;
}
