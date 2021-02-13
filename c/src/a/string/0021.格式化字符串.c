#include <stdio.h>
#include <string.h>

#define MAX 20

char *str_gets(char *str, int n);

int main(int argc, char *argv[]) {
  char first[MAX];
  char last[MAX];
  char format[2 * MAX + 10];
  double prize;
  puts("Enter your first name:");
  str_gets(first, MAX);
  puts("Enter your last name:");
  str_gets(last, MAX);
  puts("Enter your prize money:");
  scanf("%lf", &prize);
  sprintf(format, "%s, %-19s: $%6.2f\n", last, first, prize);
  puts(format);
  return 0;
}

char *str_gets(char *str, int n) {
  char *res;
  int i = 0;
  res = fgets(str, n, stdin);
  if (res) {
    while (str[i] != '\n' && str[i] != '\0') { i++; }
    if (str[i] == '\n') {
      str[i] = '\0';
    } else {
      while (getchar() != '\n') { continue; }
    }
  }
  return res;
}
