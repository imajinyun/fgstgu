#include <stdio.h>
#include <string.h>

#define SIZE 81
#define LIMIT 20
#define HALT ""

void str_sort(char *str[], int num);
char *str_gets(char *str, int n);

int main(int argc, char *argv[]) {
  char input[LIMIT][SIZE];
  char *ptr[LIMIT];
  int count = 0;
  printf("Input up to %d lines, and I will sort them.\n", LIMIT);
  printf("To stop, press the Enter key at a line's start.\n");
  while (count < LIMIT && str_gets(input[count], SIZE) && input[count][0] != '\0') {
    ptr[count] = input[count];
    count++;
  }
  str_sort(ptr, count);
  puts("\nHere's the sorted list:\n");
  for (int i = 0; i < count; i++) { puts(ptr[i]); }
  return 0;
}

void str_sort(char *str[], int num) {
  char *tmp;
  for (int i = 0; i < num - 1; i++) {
    for (int j = i + 1; j < num; j++) {
      if (strcmp(str[i], str[j]) > 0) {
        tmp = str[i];
        str[i] = str[j];
        str[j] = tmp;
      }
    }
  }
}

char *str_gets(char *str, int n) {
  char *res;
  int i = 0;
  res = fgets(str, n, stdin);
  if (res) {
    while (str[i] != '\n' && str[i] != '\0') { i++; }
    if (str[i] == '\n') {
      str[i] = '\0';
    } else {
      while (getchar() != '\n') { continue; }
    }
  }
  return res;
}
