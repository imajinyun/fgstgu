#include <stdio.h>

void showArray(const double arr[], int n);
void multArray(double arr[], int n, double mult);

void test1() {
  double rates[5] = {88.99, 77.12, 59.38, 11.29, 35.90};
  const double locked[4] = {0.08, 0.075, 0.0725, 0.07};
  const double *ptr = rates; // OK
  ptr = locked;              // OK
  ptr = &rates[3];           // OK

  printf("test1:\n");
  // 对函数的形参使用 const 不仅能保护数据，还能让函数处理 const 数组
  showArray(rates, 5);
  showArray(locked, 4);
  printf("\n");

  // 不应该把 const 数组名作为实参传递给 multArray() 这样的函数
  multArray(rates, 5, 2.5);

  // C 标准规定，使用非 const 标识符（如：multArray() 的形参 arr ）修改 const 数据（如：locked）导致的结果是未定义的
  multArray(locked, 4, 1.5); // 错误的写法
  showArray(rates, 5);
  showArray(locked, 4);
  printf("\n");
}

void test2() {
  double rates[5] = {88.99, 77.12, 59.38, 11.29, 35.90};
  const double locked[4] = {0.08, 0.075, 0.0725, 0.07};
  double *ptr = rates; // OK
  // ptr = locked;        // NO
  ptr = &rates[3]; // OK

  printf("test2:\n");
  showArray(rates, 5);
  showArray(locked, 4);
  printf("\n");
  multArray(rates, 5, 2.5);
  multArray(locked, 4, 2.5); // 错误的写法，同上
  showArray(rates, 5);
  showArray(locked, 4);
}

int main(int argc, char *argv[]) {
  test1();
  test2();
  return 0;
}

void showArray(const double arr[], int n) {
  for (int i = 0; i < n; i++) { printf("%10.3f", arr[i]); }
  putchar('\n');
}

void multArray(double arr[], int n, double mult) {
  for (int i = 0; i < n; i++) { arr[i] *= mult; }
}
