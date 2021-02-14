#include <stdio.h>

int traveler = 1;        // 全局变量，外部链接
static int stayhome = 1; // 静态变量，内部链接（具有文件作用域、内部链接和静态存储期属性）

/**
 * 内部链接的静态变量（static variable with internal linkage）。
 *
 * -> 普通的外部变量可用于同一程序中任意文件中的函数，但是内部链接的静态变量只能用于同一个文件中的函数。
 * -> 可以使用存储类别说明符extern，在函数中重复声明任何具有文件作用域的变量。这样的声明并不会改变其链接属性。
 *
 * -> C 语言中通过在一个文件中进行定义式声明，然后在其他文件中进行引用式声明来实现共享。
 *    也就是说，除了一个定义式声明外，其他声明都要使用 extern 关键字。而且，只有定义式声明才能初始化变量。
 * -> 如果外部变量定义在一个文件中，那么其他文件在使用该变量之前必须先声明它（用 extern 关键字）。
 *    也就是说，在某文件中对外部变量进行定义式声明只是单方面允许其他文件使用该变量，其他文件在用 extern 声明之前不能直接使用它。
 */
int main(int argc, char *argv[]) {
  // 对于该程序所在的翻译单元，trveler 和 stayhome 都具有文件作用域，但是只有 traveler 可用于其他翻译单元（因为它具有外部链接）。
  // 这两个声明都使用了 extern 关键字，指明了 main() 中使用的这两个变量的定义都在别处，但是这并未改变 stayhome 的内部链接属性。
  extern int traveler; // 使用定义在别处的 traveler
  extern int stayhome; // 使用定义在别处的 stayhome

  static int counter; // 具有块作用域、无链接、静态存储期属性
  return 0;
}
