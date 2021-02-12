#include <stdio.h>
#include <string.h>

#define SIZE 80

char *str_gets(char *str, int n);

int main(int argc, char *argv[]) {
  char flower[SIZE];
  char addon[] = "s smell like old shoes";
  puts("What is your favorite flower?");
  if (str_gets(flower, SIZE)) {
    strcat(flower, addon);
    puts(flower);
    puts(addon);
  } else {
    puts("End of file encountered!");
  }
  puts("Done!");
  return 0;
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
