#include <stdio.h>
#include <string.h>

size_t strnchr(const char *str, char ch) {
  int i, n = 0, len = strlen(str);
  for (i = 0; i < len; i++) {
    if (str[i] == ch) { n++; }
  }
  return n;
}

int main(int argc, char *argv[]) {
  char *str = "https://github.com/";
  char ch = '/';
  int n = strnchr(str, ch);
  printf("%c 出现的次数为：%d\n", ch, n);
  return 0;
}
