#include <stdio.h>

#define MYTYPE(x) _Generic((x), int : "int", long : "long", float : "float", double : "double", default : "other")

int main(int argc, char *argv[]) {
  int x = 5;
  printf("%s\n", MYTYPE(x));
  printf("%s\n", MYTYPE(2.0 * x));
  printf("%s\n", MYTYPE(3L));
  printf("%s\n", MYTYPE(&x));
  return 0;
}
