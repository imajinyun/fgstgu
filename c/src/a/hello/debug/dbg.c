#include "dbg.h"
#include <stdio.h>
#include <stdlib.h>

// #define NDEBUG

void testDebug() {
  debug("I have brown hair");
  debug("I am %d years old", 88);
}

void testError() {
  error("I believe everything is broken");
  error("There are %d problems in %s", 0, "space");
}

void testWarn() {
  warn("You can safely ignore this");
  warn("Maybe consider looking at: %s", "/etc/passwd");
}

void testInfo() {
  info("Well I did something mundane");
  info("It happened %f times todaay", 1.3f);
}

int testFatal(int code) {
  char *temp = (char *) malloc(100);
  checkmem(temp);
  switch (code) {
    case 1:
      info("It works");
      break;
    default:
      fatal("I shouldn't run");
  }
  free(temp);
  return 0;

goerr:
  if (temp) { free(temp); }
  return -1;
}

int testCheck(char *filename) {
  FILE *input = NULL;
  char *block = NULL;
  block = (char *) malloc(100);
  checkmem(block);

  input = fopen(filename, "r");
  check(input, "Failed to open %s", filename);
  free(block);
  fclose(input);
  return 0;

goerr:
  if (block) { free(block); }
  if (input) { fclose(input); }
  return -1;
}

int testCheckMemory() {
  char *mem = NULL;
  checkmem(mem);
  free(mem);
  return 1;

goerr:
  return -1;
}

int testCheckDebug() {
  int i = 0;
  checkdbg(i != 0, "Oops, I was 0");
  return 0;

goerr:
  return -1;
}

/**
 * 宏调试。
 *
 * -> gcc -g dbg.c -O3 -o dbg
 * -> ./dbg
 */
int main(int argc, char *argv[]) {
  testDebug();
  testError();
  testWarn();
  testInfo();

  check(testCheck("dbg.c") == 0, "failed with dbg.c");
  check(testCheck(argv[1]) == -1, "failed with argv");
  check(testFatal(1) == 0, "failed with fatal (1)");
  check(testFatal(100) == -1, "failed with fatal (100)");
  check(testCheckMemory() == -1, "failed with check memory");
  check(testCheckDebug() == -1, "failed with check debug");
  return 0;

goerr:
  return 1;
}
