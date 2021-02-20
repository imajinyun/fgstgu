#include <limits.h>
#include <stdio.h>

char *itobs(int, char *);
void show(const char *);
int invert(int num, int bit);

int main(int argc, char *argv[]) {
  char bs[CHAR_BIT * sizeof(int) + 1];
  int number;
  puts("Enter integers and see them in binary");
  puts("Non-numeric input terminates program");
  while (scanf("%d", &number) == 1) {
    itobs(number, bs);
    printf("%d is \n", number);
    show(bs);
    putchar('\n');
    number = invert(number, 4);
    printf("Inverting the last 4 bits gives\n");
    show(itobs(number, bs));
    putchar('\n');
  }
  puts("Done!");
  return 0;
}

char *itobs(int n, char *str) {
  int i;
  const static int size = CHAR_BIT * sizeof(int);
  for (i = size - 1; i >= 0; i--, n >>= 1) { str[i] = (01 & n) + '0'; }
  str[size] = '\0';
  return str;
}

void show(const char *str) {
  int i = 0;
  while (str[i]) {
    putchar(str[i]);
    if (++i % 4 == 0 && str[i]) { putchar(' '); }
  }
}

int invert(int num, int bit) {
  int mask = 0;
  int bval = 1;
  while (bit-- > 0) {
    mask |= bval;
    bval <<= 1;
  }
  return num ^ mask;
}
