#include <stdio.h>

#define ROWS 3
#define COLS 4

void sumRows(int arr[][COLS], int rows);
void sumCols(int arr[][COLS], int);
int sum(int (*arr)[COLS], int rows);

int main(int argc, char *argv[]) {
  int arr[ROWS][COLS] = {
    {2, 4, 6, 8},
    {3, 5, 7, 9},
    {12, 10, 8, 6},
  };
  sumRows(arr, ROWS);
  printf("\n");
  sumCols(arr, ROWS);
  printf("\n");
  printf("Sum of all elements: %d\n", sum(arr, ROWS));
  return 0;
}

void sumRows(int arr[][COLS], int rows) {
  int total;
  for (int i = 0; i < rows; i++) {
    total = 0;
    for (int j = 0; j < COLS; j++) { total += arr[i][j]; }
    printf("row %d: sum = %d\n", i, total);
  }
}

void sumCols(int arr[][COLS], int rows) {
  int total;
  for (int i = 0; i < COLS; i++) {
    total = 0;
    for (int j = 0; j < rows; j++) { total += arr[j][i]; }
    printf("col %d: sum = %d\n", i, total);
  }
}

int sum(int arr[][COLS], int rows) {
  int i, j, total = 0;
  for (i = 0; i < rows; i++) {
    for (j = 0; j < COLS; j++) { total += arr[i][j]; }
  }
  return total;
}
