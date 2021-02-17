#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define NLEN 81

struct names {
  char *fname;
  char *lname;
  int letters;
};

void getinfo(struct names *); // 分配内存
void makeinfo(struct names *);
void showinfo(const struct names *);
void cleanup(struct names *); // 调用该函数时释放内存
char *str_gets(char *str, int n);

int main(int argc, char *argv[]) {
  struct names person;
  getinfo(&person);
  makeinfo(&person);
  showinfo(&person);
  cleanup(&person);
  return 0;
}

void getinfo(struct names *ps) {
  char temp[NLEN];
  printf("Please enter your first name: ");
  str_gets(temp, NLEN);
  ps->fname = (char *) malloc(strlen(temp) + 1);
  strcpy(ps->fname, temp);

  printf("Please enter your last name: ");
  str_gets(temp, NLEN);
  ps->lname = (char *) malloc(strlen(temp) + 1);
  strcpy(ps->lname, temp);
  printf("\n");
}

void makeinfo(struct names *ps) {
  printf("Make person info:\n");
  ps->letters = strlen(ps->fname) + strlen(ps->lname);
  printf("%d\n\n", ps->letters);
}

void showinfo(const struct names *ps) {
  printf("Show person info:\n");
  printf("%s %s, your name contains %d letters\n", ps->fname, ps->lname, ps->letters);
}

void cleanup(struct names *ps) {
  free(ps->fname);
  free(ps->lname);
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
