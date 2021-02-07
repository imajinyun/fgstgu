#include <stdio.h>

#define ROWS 3
#define COLS 4

// 注意前两个形参（rows 和 cols）用作第 3 个形参二维数组 arr 的两个维度。
// 因为 arr 的声明要使用 rows 和 cols，所以在形参列表中必须在声明 arr 之前先声明这两个形参。
int sum2d(int rows, int cols, int arr[rows][cols]);
// int sum2d(int arr[rows][cols], int rows, int cols); // 无效的顺序

// C99/C11 标准规定，可以省略原型中的形参名，但是在这种情况下，必须用星号来代替省略的维度。
int test(int rows, int cols, int arr[*][*]);

/**
 * 变长数组（VLA - variable-length array）：
 *
 * 1. 变长数组中的『变』不是指可以修改已创建数组的大小；
 * 2. 一旦创建了变长数组，它的大小则保持不变；
 * 3. 这里的『变』指的是：在创建数组时，可以使用变量指定数组的维度；
 */
int main(int argc, char *argv[]) {
  int quarters = 4;
  int regions = 5;

  // 变长数组必须是自动存储类别，这意味着无论在函数中声明还是作为函数形参声明，
  // 都不能使用 static 或 extern 存储类别说明符。
  double sales[regions][quarters]; // 定义一个变长数组

  int rows = 3, cols = 7;
  int arr1[ROWS][COLS] = {
    {2, 4, 6, 8},
    {3, 5, 7, 9},
    {4, 5, 6, 7},
  };
  int arr2[ROWS - 1][COLS + 2] = {
    {1, 2, 3, 4, 5, 6},
    {2, 3, 4, 5, 6, 7},
  };
  int arr3[rows][cols];
  for (int i = 0; i < rows; i++) {
    for (int j = 0; j < cols; j++) { arr3[i][j] = i * j + i + j; }
  }

  printf("arr1 array:\n");
  printf("Sum of all elements = %d\n", sum2d(ROWS, COLS, arr1));
  printf("arr2 array:\n");
  printf("Sum of all elements = %d\n", sum2d(ROWS - 1, COLS + 2, arr2));
  printf("arr3 array:\n");
  printf("Sum of all elements = %d\n", sum2d(rows, cols, arr3));
  return 0;
}

int sum2d(int rows, int cols, int arr[rows][cols]) {
  int i, j, total = 0;
  for (i = 0; i < rows; i++) {
    for (j = 0; j < cols; j++) { total += arr[i][j]; }
  }
  return total;
}
