#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define SIZE 10

void showarray(const int arr[], int n);
_Static_assert(sizeof(double) == 2 * sizeof(int), "double not twice int size");

int main(int argc, char *argv[]) {
  int source[SIZE] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9};
  int target[SIZE];
  double values[SIZE / 2] = {2.0, 2.0e5, 2.0e10, 2.0e20, 5.0e30};
  puts("Source (orig of data):");
  showarray(source, SIZE);
  memcpy(target, source, SIZE * sizeof(int));
  puts("Target (copy of data):");
  showarray(target, SIZE);

  puts("\nUsing memmove() with overlapping ranges:");
  memmove(source + 2, source, 5 * sizeof(int));
  puts("Source -> element 0-4 copied to 2-6:");
  showarray(source, SIZE);

  puts("\nUsing memcpy() to copy double to int:");
  memcpy(target, values, (SIZE / 2) * sizeof(double));
  puts("Target -> 5 doubles into 10 int positions:");
  showarray(target, SIZE / 2);
  showarray(target + 5, SIZE / 2);
  return 0;
}

void showarray(const int arr[], int n) {
  for (int i = 0; i < n; i++) { printf("%d ", arr[i]); }
  putchar('\n');
}
