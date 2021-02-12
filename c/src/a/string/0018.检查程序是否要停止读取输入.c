#include <stdio.h>
#include <string.h>

#define SIZE 80
#define LIMIT 10
#define STOP "quit"

char *str_gets(char *str, int n);

int main(int argc, char *argv[]) {
  char input[LIMIT][SIZE];
  int i = 0;
  printf("Enter up to %d lines(type quit to quit):\n", LIMIT);
  // while (i < LIMIT && str_gets(input[i], SIZE) != NULL && strcmp(input[i], STOP) != 0) { i++; }
  while (i < LIMIT && str_gets(input[i], SIZE) != NULL && strcmp(input[i], STOP) != 0 && input[i][0] != '\0') { i++; }
  printf("%d strings entered\n", i);
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
