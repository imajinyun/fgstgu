#include <stdio.h>
#include <string.h>

#define SIZE 40
#define TARGSIZE 7
#define LIMIT 5

char *str_gets(char *str, int n);

int main(int argc, char *argv[]) {
  char qwords[LIMIT][TARGSIZE];
  char temp[SIZE];
  int i = 0;
  printf("Enter %d words beginning with q:\n", LIMIT);
  while (i < LIMIT && str_gets(temp, SIZE)) {
    if (temp[0] != 'q') {
      printf("%s doesn't begin with q!\n", temp);
    } else {
      // 确保存储的是一个字符串
      strncpy(qwords[i], temp, TARGSIZE - 1);
      qwords[i][TARGSIZE - 1] = '\0';
      i++;
    }
  }
  puts("Here are the words accepted:");
  for (i = 0; i < LIMIT; i++) { puts(qwords[i]); }
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
