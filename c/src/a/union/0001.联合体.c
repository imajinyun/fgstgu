#include <stdio.h>

union data {
  int x;
  char y;
  short z;
};

int main(int argc, char *argv[]) {
  union data a;
  printf("%lu %lu\n", sizeof(a), sizeof(union data));

  a.x = 0x40;
  printf("%X %c %hX\n", a.x, a.y, a.z);

  a.y = '9';
  printf("%X %c %hX\n", a.x, a.y, a.z);

  a.z = 0x2059;
  printf("%X %c %hX\n", a.x, a.y, a.z);

  a.x = 0x3E25AD54;
  printf("%X %c %hX\n", a.x, a.y, a.z);
  return 0;
}
