#include <stdio.h>
#include <stdlib.h>

void test1() {
  int n;
  int *p;
  printf("Enter array length: ");
  scanf("%d", &n);
  p = (int *) malloc(n * sizeof(int));
  int arr[n];                  // 变长数组
  p[n - 1] = arr[n - 1] = -10; // 除了索引为 n-1 的值被设定外，其它都是垃圾值
  for (int i = 0; i < n; i++) { printf("arr[%d] = %d\n", i, arr[i]); }
}

void test2() {
  int m = 6, n = 5;
  int arr[m][n]; // m*n 的变长数组（VLA）
  int(*p)[6];    // C99 之前的写法
  int(*q)[n];    // 要求支持变长数组

  p = (int(*)[6]) malloc(m * 6 * sizeof(int));     // m*6 数组
  q = (int(*)[n]) malloc(m * n * 6 * sizeof(int)); // m*n 数组（要求支持变长数组）
  arr[1][2] = p[1][2] = 33;
  printf("arr output:\n");
  for (int i = 0; i < m; i++) {
    for (int j = 0; j < n; j++) { printf("arr[%d][%d] = %d\n", i, j, arr[i][j]); }
  }
  printf("\np output:\n");
  for (int i = 0; i < m; i++) {
    for (int j = 0; j < n; j++) { printf("*p[%d][%d] = %d\n", i, j, *(*(p + i) + j)); }
  }
  printf("\nq output:\n");
  for (int i = 0; i < m; i++) {
    for (int j = 0; j < n; j++) { printf("*q[%d][%d] = %d\n", i, j, *(*(q + i) + j)); }
  }
}

int main(int argc, char *argv[]) {
  test1();
  printf("\n");
  test2();
  return 0;
}
