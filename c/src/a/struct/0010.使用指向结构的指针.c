#include <stdio.h>
#include <string.h>

#define NLEN 30

struct names {
  char fname[NLEN];
  char lname[NLEN];
  int letters;
};

void getinfo(struct names *);
void makeinfo(struct names *);
void showinfo(const struct names *);
char *str_gets(char *str, int n);

int main(int argc, char *argv[]) {
  struct names person;
  getinfo(&person);
  makeinfo(&person);
  showinfo(&person);
  return 0;
}

void getinfo(struct names *ps) {
  printf("Please enter your first name: ");
  str_gets(ps->fname, NLEN);
  printf("Please enter your last name: ");
  str_gets(ps->lname, NLEN);
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
