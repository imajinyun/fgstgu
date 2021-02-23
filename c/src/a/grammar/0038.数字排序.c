#include <stdio.h>
#include <stdlib.h>

#define NUM 40

void fillarray(double arr[], int n);
void showarray(const double arr[], int n);
int mycompare(const void *p, const void *q);

int main(int argc, char *argv[]) {
  double arr[NUM];
  fillarray(arr, NUM);
  puts("Random list:");
  showarray(arr, NUM);
  qsort(arr, NUM, sizeof(double), mycompare);
  puts("\nSorted list:");
  showarray(arr, NUM);
  return 0;
}

void fillarray(double arr[], int n) {
  for (int i = 0; i < n; i++) { arr[i] = (double) rand() / ((double) rand() + 0.1); }
}

void showarray(const double arr[], int n) {
  int i;
  for (i = 0; i < n; i++) {
    printf("%9.4f ", arr[i]);
    if (i % 6 == 5) { putchar('\n'); }
  }
  if (i % 6 != 0) { putchar('\n'); }
}

int mycompare(const void *p, const void *q) {
  const double *m = (const double *) p;
  const double *n = (const double *) q;
  if (*m > *n) {
    return 1;
  } else if (*m == *n) {
    return 0;
  } else {
    return -1;
  }
}
