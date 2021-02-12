#include <stdio.h>
#include <string.h>

void fit(char *, unsigned int);

int main(int argc, char *argv[]) {
  char msg[] = "Things should be an simple as possible,"
               " but not simpler";
  puts(msg);
  fit(msg, 38);
  puts(msg);
  puts("Let's look at some more of the string");
  puts(msg + 39);
  return 0;
}

void fit(char *str, unsigned int size) {
  if (strlen(str) > size) { str[size] = '\0'; }
}
