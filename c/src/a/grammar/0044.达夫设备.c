#include "../hello/debug/dbg.h"
#include <stdio.h>
#include <string.h>

int normalCopy(char *from, char *to, int n) {
  int i;
  for (i = 0; i < n; i++) { to[i] = from[i]; }
  return i;
}

int duffsDevice(char *from, char *to, int n) {
  {
    int m = (n + 7) / 8;
    switch (n % 8) {
      case 0:
        do {
          *to++ = *from++;
          case 7:
            *to++ = *from++;
          case 6:
            *to++ = *from++;
          case 5:
            *to++ = *from++;
          case 4:
            *to++ = *from++;
          case 3:
            *to++ = *from++;
          case 2:
            *to++ = *from++;
          case 1:
            *to++ = *from++;
        } while (--m > 0);
    }
  }
  return n;
}

int validCopy(char *data, int n, char expects) {
  int i;
  for (i = 0; i < n; i++) {
    if (data[i] != expects) {
      error("[%d] %c != %c", i, data[i], expects);
      return 0;
    }
  }
  return 1;
}

int main(int argc, char *argv[]) {
  char from[1000] = {'a'};
  char to[1000] = {'c'};
  int row = 0;

  memset(from, 'x', 1000);
  memset(to, 'y', 1000);

  check(validCopy(to, 1000, 'y'), "Not initialized right");

  // use normal copy to
  row = normalCopy(from, to, 1000);
  check(row == 1000, "Normal copy failed: %d", row);
  check(validCopy(to, 1000, 'x'), "Normal copy failed");

  // reset
  memset(to, 'y', 1000);

  // duffs version
  row = duffsDevice(from, to, 1000);
  check(row == 1000, "Duff's device failed: %d", row);
  check(validCopy(to, 1000, 'x'), "Duff's device failed copy'");

  memset(to, 'y', 1000);
  debug("This program no output");
  return 0;

goerr:
  return 1;
}
