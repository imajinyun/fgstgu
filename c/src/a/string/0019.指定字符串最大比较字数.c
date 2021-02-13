#include <stdio.h>
#include <string.h>

#define SIZE 6

char *str_gets(char *str, int n);

int main(int argc, char *argv[]) {
  const char *list[SIZE] = {
    "astronomy", "astounding", "astrophysics", "ostracize", "asterism", "astrophobia",
  };
  int count = 0;
  int i;
  for (i = 0; i < SIZE; i++) {
    if (strncmp(list[i], "astro", 5) == 0) {
      printf("Found: %s\n", list[i]);
      count++;
    }
  }
  printf("The list contained %d words beginning"
         " with astro\n",
         count);
  return 0;
}

char *str_gets(char *str, int n) {
  char *res;
  int i = 0;
  res = fgets(str, n, stdin);
  if (res) {
    while (str[i] != '\0' && str[i] != '\n') { i++; }
    if (str[i] == '\n') {
      str[i] = '\0';
    } else {
      while (getchar() != '\0') { continue; }
    }
  }
  return res;
}
