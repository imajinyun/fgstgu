#include <stdio.h>
#include <stdlib.h>

struct flex {
  size_t size;
  double average;
  double scores[]; // 伸缩型数组成员
};

void showflex(const struct flex *p);

int main(int argc, char *argv[]) {
  struct flex *f1, *f2;
  int n = 5;
  int i;
  int total = 0;
  f1 = malloc(sizeof(struct flex) + n * sizeof(double));
  f1->size = n;
  for (i = 0; i < n; i++) {
    f1->scores[i] = 20.0 - i;
    total += f1->scores[i];
  }
  f1->average = total / n;
  showflex(f1);

  n = 9;
  total = 0;
  f2 = malloc(sizeof(struct flex) + n * sizeof(double));
  f2->size = n;
  for (i = 0; i < n; i++) {
    f2->scores[i] = 20.0 - i / 2.0;
    total += f2->scores[i];
  }
  f2->average = total / n;
  showflex(f2);
  free(f1);
  free(f2);
  return 0;
}

void showflex(const struct flex *p) {
  int i;
  printf("Scores: ");
  for (i = 0; i < p->size; i++) { printf("%g ", p->scores[i]); }
  printf("\nAverage: %g\n", p->average);
}
