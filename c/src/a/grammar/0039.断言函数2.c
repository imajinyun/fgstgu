#include <limits.h>
#include <stdio.h>

// -> error: static_assert failed due to requirement '8 == 16' "16-bit char falsely assumed"
_Static_assert(CHAR_BIT == 16, "16-bit char falsely assumed");

int main(int argc, char *argv[]) {
  puts("char 16 is bits");
  return 0;
}
