#include "name_st.h"
#include <stdio.h>

void get_names(names *p) {
  printf("Please enter your first name: ");
  str_get(p->first, SLEN);
  printf("Please enter your last name: ");
  str_get(p->last, SLEN);
}

void show_names(const names *p) { printf("%s %s", p->first, p->last); }

char *str_get(char *str, int n) {
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
