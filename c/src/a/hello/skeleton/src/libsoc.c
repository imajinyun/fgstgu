#include "dbg.h"
#include <ctype.h>
#include <stdio.h>

int message(const char *msg) {
  printf("A string: %s\n", msg);
  return 0;
}

int uppercase(const char *msg) {
  int i;
  for (i = 0; msg[i] != '\0'; i++) { printf("%c", toupper(msg[i])); }
  printf("\n");
  return 0;
}

int lowercase(const char *msg) {
  int i;
  for (i = 0; msg[i] != '\0'; i++) { printf("%c", tolower(msg[i])); }
  printf("\n");
  return 0;
}

int failure(const char *msg) { return 1; }
