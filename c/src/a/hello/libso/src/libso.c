#include "dbg.h"
#include <ctype.h>
#include <stdio.h>

int message(const char *msg) {
  printf("You have a message: %s\n", msg);
  return 0;
}

int failure(const char *msg) {
  printf("You have a failure: %s\n", msg);
  return 1;
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
