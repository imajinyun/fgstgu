#include <stdbool.h>
#include <stdio.h>
#include <string.h>

bool findCharInString(char *source, char target);

int main(int argc, char const *argv[]) {
  char letter;
  char url[] = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
  printf("Please input a char: ");
  letter = getchar();

  if (findCharInString(url, letter)) {
    printf("%c exists in the given string\n", letter);
  } else {
    printf("%c not exists in the given string\n", letter);
  }
  return 0;
}

bool findCharInString(char *source, char target) {
  for (int i = 0, n = strlen(source); i < n; i++) {
    if (source[i] == target) { return true; }
  }
  return false;
}
