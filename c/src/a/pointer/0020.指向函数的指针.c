#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef int (*compare)(int a, int b);

void die(const char *message) {
  if (errno) {
    perror(message);
  } else {
    printf("ERROR: %s\n", message);
  }
  exit(EXIT_FAILURE);
}

int *bubbleSort(int *nums, int n, compare cmp) {
  int tmp = 0, i, j;
  int *dst = malloc(n * sizeof(int));
  if (!dst) { die("Memory allocation failed"); }
  memcpy(dst, nums, n * sizeof(int));
  for (i = 0; i < n; i++) {
    for (j = 0; j < n - 1; j++) {
      if (cmp(dst[j], dst[j + 1]) > 0) {
        tmp = dst[j];
        dst[j] = dst[j + 1];
        dst[j + 1] = tmp;
      }
    }
  }
  return dst;
}

int sortedOrder(int a, int b) { return a - b; }

int reverseOrder(int a, int b) { return b - a; }

int strangeOrder(int a, int b) {
  if (a == 0 || b == 0) {
    return 0;
  } else {
    return a % b;
  }
}

void testSorting(int *nums, int n, compare cmp) {
  int i;
  int *sorted = bubbleSort(nums, n, cmp);
  if (!sorted) { die("Failed to sort as requested"); }
  for (i = 0; i < n; i++) { printf("%d ", sorted[i]); }
  printf("\n");
  free(sorted);
}

int main(int argc, char *argv[]) {
  if (argc < 2) { die("Usage: /path/to/a.pointer.0020 4 1 0 5 2"); }
  int n = argc - 1;
  int i = 0;
  char **input = argv + 1;
  int *nums = malloc(n * sizeof(int));
  if (!nums) { die("Memory allocation failed"); }
  for (i = 0; i < n; i++) { nums[i] = atoi(input[i]); }
  testSorting(nums, n, sortedOrder);
  testSorting(nums, n, reverseOrder);
  testSorting(nums, n, strangeOrder);
  free(nums);
  return 0;
}
