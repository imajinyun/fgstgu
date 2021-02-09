#include <stdio.h>

#define STRLEN 10

int main(int argc, char *argv[]) {
  char words[STRLEN];
  puts("Enter strings(empty line to quit):");
  while (fgets(words, STRLEN, stdin) != NULL && words[0] != '\n') { fputs(words, stdout); }
  puts("Done!");
  return 0;
}
