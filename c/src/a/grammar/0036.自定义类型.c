#include <stdio.h>

#define BYTE_DEF unsigned char

typedef unsigned char BYTE_TYP;
typedef unsigned char byte;

// 没有 typedef 关键字，编译器将把 STRING 识别为一个指向 char 的指针变量。
// 有了 typedef 关键字，编译器则把 STRING 解释成一个类型的标识符，该类型是指向 char 的指针。
typedef char *STRING;

// 可使用 COMPLEX 类型代替 complex 结构来表示复数。
typedef struct complex {
  float real;
  float imag;
} COMPLEX;

// 用 typedef 来命名一个结构类型时，可以省略该结构的标签。
typedef struct {
  double x;
  double y;
} rect;

int main(int argc, char *argv[]) {
  puts("typedef keyword concept!\n");
  BYTE_TYP x, y[10], *z;
  byte o, p[10], *q;

  STRING name, sign; // 相当于 char *name, *sign;

  puts("使用 typedef 的第 1 个原因是：为经常出现的类型创建一个方便、易识别的类型名。\n");
  rect r1 = {3.0, 6.0}; // 相当于 struct {double x; double y;} r1 = {3.0, 6.0};
  rect r2;              // 相当于 struct {double x; double y;} r2;
  r2 = r1; // 这两个结构在声明时都没有标记，它们的成员完全相同，所以 r1 和 r2 间的赋值是有效操作

  puts("使用 typedef 的第 2 个原因是：typedef 常用于给复杂的类型命名。");
  typedef char(*FRPTC())[5]; // 把 FRPTC 声明为一个函数类型，该函数返回一个指针，该指针指向内含 5 个 char 类型元素的数组
  return 0;
}
