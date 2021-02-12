#include <ctype.h>
#include <stdio.h>
#include <string.h>

#define ANSWER "grant"
#define SIZE 40

char *str_gets(char *str, int n);
char *str_lower(char *str);

int main(int argc, char *argv[]) {
  char tries[SIZE];
  puts("Who is buried in Grant's tomb?");
  str_gets(tries, SIZE);
  while (strcmp(str_lower(tries), ANSWER)) {
    puts("No, that's wrong, try again");
    str_gets(tries, SIZE);
  }
  puts("That's right!");
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

char *str_lower(char *str) {
  char *res = str;
  for (; *str != '\0'; str++) { *str = tolower(*str); }
  return res;
}
