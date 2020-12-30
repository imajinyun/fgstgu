#include <stdio.h>
#include <stdlib.h>

#define DIM(x) (sizeof(x) / sizeof((x)[0]))

static int compare(const void *a, const void *b) {
  const int *pa = (int *)a;
  const int *pb = (int *)b;
  return *pa - *pb;
}

// -> gcc -o qsort.o qsort.c
// -> ./qsort.o
int main(int argc, const char *argv[]) {
  int values[] = {10, 9, 1, 0, 2, 101, 98, 17, 6, 33, 78};
  int i;
  qsort(values, DIM(values), sizeof(values[0]), compare);
  for (i = 0; i < DIM(values); i++) {
    printf("%d", values[i]);
  }
  return 0;
}
