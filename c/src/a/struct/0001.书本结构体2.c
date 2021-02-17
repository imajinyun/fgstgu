#include <stdio.h>
#include <string.h>

#define MAX_TITLE 41  // 书名的最大长度 + 1
#define MAX_AUTHOR 31 // 作者姓名的最大长度 + 1
#define MAX_BOOK 101

struct book {
  char title[MAX_TITLE];
  char author[MAX_AUTHOR];
  float value;
};

char *str_gets(char *str, int n);

int main(int argc, char *argv[]) {
  struct book library[MAX_BOOK];
  int count = 0;
  int index;
  printf("Please enter the book title ([Enter] at the start of a line to stop): ");
  while (count < MAX_BOOK && str_gets(library[count].title, MAX_TITLE) != NULL && library[count].title[0] != '\0') {
    printf("Now enter the author: ");
    str_gets(library[count].author, MAX_AUTHOR);
    printf("Now enter the value: ");
    scanf("%f", &library[count++].value);
    while (getchar() != '\n') { continue; }
    if (count < MAX_BOOK) { printf("Enter the next title: "); }
  }
  if (count > 0) {
    printf("Here is the list of your books:\n");
    for (index = 0; index < count; index++) {
      printf("%s by %s: $%.2f\n", library[index].title, library[index].author, library[index].value);
    }
  } else {
    printf("No books? Too bad!\n");
  }
  return 0;
}

char *str_gets(char *str, int n) {
  char *res, *find;
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
