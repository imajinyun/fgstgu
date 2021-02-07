#include <stdio.h>

#define COLS 4

int sum(const int arr[], int n);
int sum2d(const int arr[][COLS], int rows);

void test1() {
  int a[2] = {10, 20};
  printf("a[0] = %d\n", a[0]);
  printf("a[1] = %d\n", a[1]);
  (int[2]){10, 20};
  (int[]){11, 22, 33, 44, 55};
  printf("(int[2]){11, 22}[0] = %d\n", (int[2]){11, 22}[0]);

  int *p;
  p = (int[2]){1, 2};
  printf("*p = %d\n", *p);
  printf("*(p + 1) = %d\n", *(p + 1));
}

void test2() {
  int t1, t2, t3;
  int *p1;
  int(*p2)[COLS];
  p1 = (int[2]){10, 20};
  p2 = (int[2][COLS]){{1, 2, 3, 4}, {5, 6, 7, 8}};

  t1 = sum(p1, 2);
  t2 = sum2d(p2, 2);
  t3 = sum((int[]){1, 2, 3, 4, 5, 6, 7, 8, 9}, 9);

  printf("\n");
  printf("t1 = %d\n", t1);
  printf("t2 = %d\n", t2);
  printf("t3 = %d\n", t3);
}

/**
 * 复合字面量：
 * 1. 提供只临时需要的值的一种手段；
 * 2. 复合字面量具有块作用域，这意味着一旦离开定义复合字面量的块，程序将无法保证该字面量是否存在；
 * 3. 也就是说，复合字面量的定义在最内层的花括号中；
 */
int main(int argc, char *argv[]) {
  test1();
  test2();
  return 0;
}

int sum(const int arr[], int n) {
  int i, total = 0;
  for (i = 0; i < n; i++) { total += arr[i]; }
  return total;
}

int sum2d(const int arr[][COLS], int rows) {
  int i, j, total = 0;
  for (i = 0; i < rows; i++) {
    for (j = 0; j < COLS; j++) { total += arr[i][j]; }
  }
  return total;
}
