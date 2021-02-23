#include <stdio.h>
#include <string.h>

#define TSIZE 45
#define FMAX 5

struct film {
  char title[TSIZE];
  int rating;
};

char *str_gets(char *str, int n);

int main(int argc, char *argv[]) {
  struct film movies[FMAX];
  int i = 0, j;
  puts("Enter first movie title:");
  while (i < FMAX && str_gets(movies[i].title, TSIZE) != NULL && movies[i].title[0] != '\0') {
    puts("Enter your rating (0-10):");
    scanf("%d", &movies[i++].rating);
    while (getchar() != '\n') { continue; }
    puts("Enter next movie title (empty line to stop):");
  }
  if (i == 0) {
    printf("No data entered");
  } else {
    printf("Here is the movie list:\n");
    for (j = 0; j < i; j++) { printf("Movie: %s\t\tRating: %d\n", movies[j].title, movies[j].rating); }
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
