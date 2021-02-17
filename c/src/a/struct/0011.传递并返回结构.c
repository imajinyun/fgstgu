#include <stdio.h>
#include <string.h>

#define NLEN 30

struct names {
  char fname[NLEN];
  char lname[NLEN];
  int letters;
};

struct names getinfo();
struct names makeinfo(struct names);
void showinfo(struct names);
char *str_gets(char *str, int n);

int main(int argc, char *argv[]) {
  struct names person;
  person = getinfo();
  person = makeinfo(person);
  showinfo(person);
  return 0;
}

struct names getinfo() {
  struct names person;
  printf("Please enter your first name: ");
  str_gets(person.fname, NLEN);
  printf("Please enter your last name: ");
  str_gets(person.lname, NLEN);
  return person;
}

struct names makeinfo(struct names person) {
  person.letters = strlen(person.fname) + strlen(person.lname);
  return person;
}

void showinfo(struct names person) {
  printf("\nShow person info:\n");
  printf("%s %s, your name contains %d letters\n", person.fname, person.lname, person.letters);
}

char *str_gets(char *str, int n) {
  char *res, *find;
  res = fgets(str, n, stdin);
  if (res) {
    if ((find = strchr(str, '\n')) != NULL) {
      *find = '\0';
    } else {
      while (getchar() != '\n') { continue;}
    }
  }
  return res;
}
