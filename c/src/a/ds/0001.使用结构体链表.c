#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define TSIZE 45

struct film {
  char title[TSIZE];
  int rating;
  struct film *next;
};

char *str_gets(char *str, int n);

int main(int argc, char *argv[]) {
  struct film *head = NULL;
  struct film *prev, *curr;
  char input[TSIZE];
  puts("Enter first movie title:");
  while (str_gets(input, TSIZE) != NULL && input[0] != '\0') {
    curr = (struct film *) malloc(sizeof(struct film));
    if (head == NULL) {
      head = curr;
    } else {
      prev->next = curr;
    }
    curr->next = NULL;
    strcpy(curr->title, input);
    puts("Enter your rating (0~10):");
    scanf("%d", &curr->rating);
    while (getchar() != '\n') { continue; }
    puts("Enter next movie title (empty line to stop):");
    prev = curr;
  }
  if (head == NULL) {
    printf("No data entered.\n");
  } else {
    printf("Here is the movie list:\n");
    curr = head;
    while (curr != NULL) {
      printf("Movie: %s\t\tRating: %d\n", curr->title, curr->rating);
      curr = curr->next;
    }
  }
  curr = head;
  while (curr != NULL) {
    head = curr->next;
    free(curr);
    curr = head;
  }
  printf("Done!\n");
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
