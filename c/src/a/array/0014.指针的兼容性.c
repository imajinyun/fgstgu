#include <stdio.h>

void test1() {
  int n = 5;
  double x;
  int *p1 = &n;
  double *p2 = &x;
  x = n; // 隐式类型转换

  // -> incompatible pointer types assigning to 'double *' from 'int *'
  // p2 = p1; // 编译时错误
}

void test2() {
  int *pt;
  int(*pa)[3];
  int arr1[2][3];
  int arr2[3][2];
  int **p2; // 一个指向指针的指针

  pt = &arr1[0][0]; // 都是指向 int 的指针
  pt = arr1[0];     // 都是指向 int 的指针
  pt = arr1;        // 无效
  pa = arr1;        // 都是指向包含 3 个 int 类型元素数组的指针
  pa = arr2;        // 无效
  p2 = &pt;         // 都是指向 int * 的指针
  *p2 = arr2[0];    // 都是指向 int 的指针
  p2 = arr2;        // 无效
}

void test3() {
  int x = 20;
  const int y = 99;
  int *p1 = &x;
  const int *p2 = &y;
  const int **pp2;

  p1 = p2;   // 不安全，把 const 指针赋值给非 const 指针
  p2 = p1;   // 有效，把非 const 指针赋给 const 指针
  pp2 = &p1; // 不安全，嵌套指针类型赋值
}

void test4() {
  const int **pp2;
  int *p1;
  const int n = 99;

  pp2 = &p1; // 允许，但是这导致 const 限定符失效（但是根据第 1 行代码，不能通过 **pp2 修改它所指向的内容）
  *pp2 = &n; // 有效，两者都声明为 const，但是这将导致 p1 指向 n（*pp2 已经被修改）
  *p1 = 10;  // 有效，但是这将改变 n 的值（但是根据第 3 行代码，不能修改 n 的值）
}

void test5() {
  /**
   * C 和 C++ 中 const 的用法区别：
   * 1. C++ 允许在声明数组大小时使用 const 整数，而 C 却不允许；
   * 2. C++ 的指针赋值检查更严格；
   */
  const int y;
  const int *p2 = &y;
  int *p1;
  p1 = p2; // C++ 中不允许这样做，但是 C 可能只给警告
}

int main(int argc, char *argv[]) {
  test1();
  test2();
  test3();
  test4();
  test5();
  return 0;
}
