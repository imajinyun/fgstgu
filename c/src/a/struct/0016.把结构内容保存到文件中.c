#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_TITLE 41
#define MAX_AUTHOR 41
#define MAX_BOOK 10

char *str_gets(char *str, int n);

struct mybook {
  char title[MAX_TITLE];
  char author[MAX_TITLE];
  float value;
};

int main(int argc, char *argv[]) {
  struct mybook book[MAX_BOOK];
  int count = 0;
  int index, nfile;
  FILE *pb;
  int size = sizeof(struct mybook);
  if ((pb = fopen("/tmp/book.dat", "a+b")) == NULL) {
    fputs("Can't open book.dat file\n", stderr);
    exit(1);
  }
  rewind(pb);
  while (count < MAX_BOOK && fread(&book[count], size, 1, pb) == 1) {
    if (count == 0) {
      puts("Current contents of book.dat:");
      printf("%s by %s: ￥%.2f\n", book[count].title, book[count].author, book[count].value);
    }
    count++;
  }
  nfile = count;
  if (count == MAX_BOOK) {
    fputs("The book.dat file is full\n", stderr);
    exit(2);
  }
  puts("Please add new book titles (press [enter] at the start of a line to stop)");
  while (count < MAX_BOOK && str_gets(book[count].title, MAX_TITLE) != NULL && book[count].title[0] != '\0') {
    puts("Now enter the author: ");
    str_gets(book[count].author, MAX_AUTHOR);
    puts("Now enter the value: ");
    scanf("%f", &book[count++].value);
    while (getchar() != '\n') { continue; }
    if (count < MAX_BOOK) { puts("Enter the next title"); }
  }

  if (count > 0) {
    puts("Here is the list of your books:");
    for (index = 0; index < count; index++) {
      printf("%s by %s: ￥%.2f\n", book[index].title, book[index].author, book[index].value);
    }
    fwrite(&book[nfile], size, count - nfile, pb);
  } else {
    puts("No books? Too bad\n");
  }
  puts("Done!");
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
