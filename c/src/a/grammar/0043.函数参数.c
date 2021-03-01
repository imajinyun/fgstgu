#include <ctype.h>
#include <stdio.h>

int canPrintIt(char ch);
void printLetters(char argv[]);
void printArgs(int argc, char *argv[]);

int main(int argc, char *argv[]) {
  printArgs(argc, argv);
  return 0;
}

int canPrintIt(char ch) { return isalpha(ch) || isblank(ch); }

void printLetters(char argv[]) {
  int i = 0;
  for (i = 0; argv[i] != '\0'; i++) {
    char ch = argv[i];
    if (canPrintIt(ch)) { printf("%c=%d ", ch, ch); }
  }
  printf("\n");
}

void printArgs(int argc, char *argv[]) {
  // 设置 i = 1 试试
  for (int i = 0; i < argc; i++) { printLetters(argv[i]); }
}
