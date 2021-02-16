#include <stdio.h>
#include <string.h>

#define MAX_TITLE 41  // 书名的最大长度 + 1
#define MAX_AUTHOR 31 // 作者姓名的最大长度 + 1

char *str_gets(char *str, int n);

struct book {
  char title[MAX_TITLE];
  char author[MAX_AUTHOR];
  float value;
};

int main(int argc, char *argv[]) {
  struct book library;
  printf("Please enter the book title: ");
  str_gets(library.title, MAX_TITLE);
  printf("Now enter the author: ");
  str_gets(library.author, MAX_AUTHOR);
  printf("Now enter the value: ");
  scanf("%f", &library.value);
  printf("%s by %s: $%.2f\n", library.title, library.author, library.value);
  printf("Done!\n");
  return 0;
}

char *str_gets(char *str, int n) {
  char *res;
  char *find;
  res = fgets(str, n, stdin);
  if (res) {
    if ((find = strchr(str, '\n')) != NULL) {
      *find = '\0';
    } else {
      while (getchar() != '\n') { continue; }
    }
  }
  return res;
}
