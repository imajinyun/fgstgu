#undef NDEBUG
#ifndef MINIUNIT_H__
#define MINIUNIT_H__

#include "dbg.h"
#include <stdio.h>
#include <stdlib.h>

#define TestStart() char *message = NULL

#define Assert(test, message)                                                  \
  if (!(test)) {                                                               \
    error(message);                                                            \
    return message;                                                            \
  }

#define RunTestCase(test)                                                      \
  debug("\n=== RUN    %s", " " #test);                                         \
  message = test();                                                            \
  tcounter++;                                                                  \
  if (message) {                                                               \
    printf("\n--- FAIL   %s", " " #test);                                      \
    return message;                                                            \
  } else {                                                                     \
    printf("\n--- PASS   %s", " " #test);                                      \
  }

#define RunTestMain(name)                                                      \
  int main(int argc, char *argv[]) {                                           \
    argc = 1;                                                                  \
    debug("\n=== RUN:    %s", argv[0]);                                        \
    printf("\n=== RUN:    %s\n", argv[0]);                                     \
    char *res = name();                                                        \
    if (res != 0) {                                                            \
      printf("\n\n--- FAIL:   %s\n", res);                                     \
    } else {                                                                   \
      printf("\n\n--- STATUS: ✔️\n");                                      \
    }                                                                          \
    printf("--- TOTAL:  %d\n", tcounter);                                      \
    exit(res != 0);                                                            \
  }

int tcounter;

#endif
