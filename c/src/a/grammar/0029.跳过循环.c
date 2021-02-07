#include <stdio.h>

int main(int argc, char *argv[]) {
  char ch = 0;

  printf("跳过循环：\n");
  while (ch != '\n') {
    ch = getchar();

    if (ch == 'a' || ch == '1' || ch == 'z' || ch == '9') { continue; }
    putchar(ch);
    return 0;
  }
}
