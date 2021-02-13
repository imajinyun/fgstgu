#include <stdio.h>
#include <string.h>

#define WORDS "beast"
#define SIZE 40

int main(int argc, char *argv[]) {
  // strcpy() 和 strcat() 都有同样的问题，它们都不能检查目标空间是否能容纳源字符串的副本。
  const char *orig = WORDS;
  char copy[SIZE] = "Be the best that you can be";
  char *ps;
  puts(orig);
  puts(copy);
  ps = strcpy(copy + 7, orig);
  puts(orig);
  puts(copy);
  puts(ps);
  return 0;
}
