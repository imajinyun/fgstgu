#include <stdio.h>
#include <stdlib.h>

void mallocAfterFree(double arr[], int n);
void callocAfterFree(double arr[], int n);
void testMalloc();
void testCalloc();

int main(int argc, char *argv[]) {
  testMalloc();
  printf("\n");
  testCalloc();
  return 0;
}

void mallocAfterFree(double arr[], int n) {
  // 如果忘记使用 free 释放，可能会存在内存泄漏（memory leak）的情况
  double *temp = (double *) malloc(n * sizeof(double));
  free(temp);
}

void callocAfterFree(double arr[], int n) {
  // calloc() 函数把块中的所有位都设置为 0（注意，在某些硬件系统中，不是把所有位都设置为 0 来表示浮点值 0）
  double *temp = (double *) calloc(n, sizeof(double));
  free(temp);
}

void testMalloc() {
  double arr[2000];
  printf("释放 malloc() 分配的内存：\n");
  printf("循环要执行 1000 次，所以在循环结束时，内存池中有 1600 万字节被占用\n");
  for (int i = 0; i < 1000; i++) { mallocAfterFree(arr, 2000); }
}

void testCalloc() {
  double arr[2000];
  printf("释放 calloc() 分配的内存：\n");
  printf("循环要执行 1000 次，所以在循环结束时，内存池中有 1600 万字节被占用\n");
  for (int i = 0; i < 1000; i++) { callocAfterFree(arr, 2000); }
}
