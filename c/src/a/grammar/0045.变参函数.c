#include "../hello/debug/dbg.h"
#include <stdarg.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_DATA 100

int readString(char **out, int buf) {
  *out = calloc(1, buf + 1);
  checkmem(*out);

  char *res = fgets(*out, buf, stdin);
  check(res != NULL, "input error");
  return 0;
goerr:
  if (*out) { free(*out); }
  *out = NULL;
  return -1;
}

int readInt(long *out) {
  char *arg = NULL, *end = NULL;
  int n = readString(&arg, MAX_DATA);
  check(n == 0, "failed to read number");

  *out = strtol(arg, &end, 10);
  bool isend = *end == '\0' || *end == '\n';
  check(isend && *arg != '\0', "invalid number: %s", arg);
  free(arg);
  return 0;
goerr:
  if (arg) { free(arg); }
  return -1;
}

int readScan(const char *fmt, ...) {
  int i, n = 0;
  long *num = NULL;
  char *chr = NULL;
  char **str = NULL;
  int buf = 0;

  va_list ap;
  va_start(ap, fmt);
  for (i = 0; fmt[i] != '\0'; i++) {
    if (fmt[i] == '%') {
      i++;
      switch (fmt[i]) {
        case '\0':
          fatal("invalid format, you ended with %%");
          break;
        case 'd':
          num = va_arg(ap, long *);
          n = readInt(num);
          check(n == 0, "failed to read int");
          break;
        case 'c':
          chr = va_arg(ap, char *);
          *chr = fgetc(stdin);
          break;
        case 's':
          buf = va_arg(ap, int);
          str = va_arg(ap, char **);
          n = readString(str, buf);
          check(n == 0, "failed to read string");
          break;
        default:
          fatal("invalid format");
      }
    } else {
      fgetc(stdin);
    }
    check(!feof(stdin) && !ferror(stdin), "input error");
  }
  va_end(ap);
  return 0;

goerr:
  va_end(ap);
  return -1;
}

int main(int argc, char *argv[]) {
  char *fname = NULL, *lname = NULL;
  char initial = ' ';
  long age = 0;

  printf("What's your first name? ");
  int n = readScan("%s", MAX_DATA, &fname);
  check(n == 0, "failed first name");

  printf("What's your initial? ");
  n = readScan("%c\n", &initial);
  check(n == 0, "failed initial");

  printf("What's your last name? ");
  n = readScan("%s", MAX_DATA, &lname);
  check(n == 0, "failed last name");

  printf("How old are you? ");
  n = readScan("%d", &age);
  check(n == 0, "failed to read age");

  printf("\nOutput:\n");
  printf("  First Name: %s", fname);
  printf("  Initial: '%c'\n", initial);
  printf("  Last Name: %s", lname);
  printf("  Age: %ld\n", age);

  free(fname);
  free(lname);
  return 0;

goerr:
  return -1;
}
