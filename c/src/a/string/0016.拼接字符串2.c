#include <stdio.h>
#include <string.h>

#define SIZE 30
#define BUGSIZE 13

char *str_gets(char *, int);

int main(int argc, char *argv[]) {
  // strcat() 和 gets() 类似，也会导致缓冲区溢出。
  // gets() 造成的安全隐患来自于使用该程序的人，而 strcat() 暴露的问题是由程序员造成的。
  char flower[SIZE];
  char addon[] = "s smell like old shose";
  char bug[BUGSIZE];
  int available;
  puts("What is your favorite flower?");
  str_gets(flower, SIZE);
  if ((strlen(addon) + strlen(flower) + 1) <= SIZE) { strcat(flower, addon); }
  puts(flower);
  puts("What is your favorite bug?");
  str_gets(bug, BUGSIZE);
  available = BUGSIZE - strlen(bug) - 1;
  strncat(bug, addon, available);
  puts(bug);
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
