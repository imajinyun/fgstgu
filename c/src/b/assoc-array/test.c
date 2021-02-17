#include <stdio.h>
#include <string.h>

#include "hash_table.h"

static char *all_tests() { return 0; }

int main(int argc, char *argv[]) {
  char *res = all_tests();
  if (res != 0) {
    printf("%s\n", res);
  } else {
    printf("all test passed\n");
  }
  return 0;
}
