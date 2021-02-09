#include <stdio.h>

#define STRLEN 81

char *str_gets(char *str, int n);

int main(int argc, char *argv[]) {
  char words[STRLEN];
  puts("Enter strings(empty line to quit):");
  puts(str_gets(words, 5));
  puts("Done!");
  return 0;
}

char *str_gets(char *str, int n) {
  char *res;
  int i = 0;
  res = fgets(str, n, stdin);
  if (res != NULL) {
    while (str[i] != '\n' && str[i] != '\0') { i++; }
    if (str[i] == '\n') {
      str[i] = '\0';
    } else {
      while (getchar() != '\n') { continue; }
    }
  }
  return res;
}
