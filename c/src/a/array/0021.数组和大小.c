#include <stdio.h>

int main(int argc, char *argv[]) {
  int areas[] = {10, 12, 13, 14, 20};
  char sname[] = "Hello";
  char fname[] = {'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '\0'};

  printf("The size of an int: %ld\n", sizeof(int));
  printf("The size of areas (int[]): %ld\n", sizeof(areas));
  printf("The number of ints in areas: %ld\n", sizeof(areas) / sizeof(int));
  printf("The first area is %d, the second is %d\n", areas[0], areas[1]);
  printf("\n");
  printf("The size of a char: %ld\n", sizeof(char));
  printf("The size of sname (char[]): %ld\n", sizeof(sname));
  printf("The number of chars: %ld\n", sizeof(sname) / sizeof(char));
  printf("The size of fname (char[]: %ld\n", sizeof(fname));
  printf("The number of chars: %ld\n", sizeof(fname) / sizeof(char));
  printf("short name is %s, full name is %s\n", sname, fname);
  return 0;
}
