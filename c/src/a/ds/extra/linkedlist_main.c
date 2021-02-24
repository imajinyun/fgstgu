#include "linkedlist.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void showmovies(Item item);
char *str_gets(char *str, int n);

int main(int argc, char *argv[]) {
  List movies;
  Item curr;

  InitializeList(&movies);
  if (IsFullList(&movies)) {
    fprintf(stderr, "No memory available!\n");
    exit(EXIT_FAILURE);
  }
  puts("Enter first movie title:");
  while (str_gets(curr.title, TSIZE) != NULL && curr.title[0] != '\0') {
    puts("Enter your rating (0~10):");
    scanf("%d", &curr.rating);
    while (getchar() != '\n') { continue; }
    if (ListAddItem(curr, &movies) == false) {
      fprintf(stderr, "Problem allocating memory!\n");
      break;
    }
    if (IsFullList(&movies)) {
      puts("The list is full now!\n");
      break;
    }
    puts("Enter next movie title (empty line to stop):");
  }

  if (IsEmptyList(&movies)) {
    printf("No data entered!\n");
  } else {
    printf("Here is the movie list:\n");
    ListTraverse(&movies, showmovies);
  }
  printf("You entered %d movies.\n", ListItemCount(&movies));
  ListRelease(&movies);
  puts("Done!");
  return 0;
}

void showmovies(Item item) { printf("Movie: %s\t\tRating: %d\n", item.title, item.rating); }

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
