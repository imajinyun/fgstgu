#include <stdio.h>

int main(int argc, char *argv[]) {
  char str[20] = {0};
  int i;

  for (i = 0; i < 10; i++) { *(str + i) = 97 + i; }

  printf("       str = %s\n", str);          // abcdefghij
  printf("     str+1 = %s\n", str + 1);      // bcdefghij
  printf("    str[2] = %c\n", str[2]);       // c
  printf("(str+2)[2] = %c\n", (str + 2)[2]); // e
  return 0;
}
