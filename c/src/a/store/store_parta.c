#include <stdio.h>

void report_count();
void accumulate(int k);

int count = 0; // 文件作用域，外部链接

/*
 * 存储类别：
 *
 * 1. 自动变量具有块作用域、无链接、自动存储期。它们是局部变量，属于其定义所在块（通常指函数）私有。
 *    寄存器变量的属性和自动变量相同，但是编译器会使用更快的内存或寄存器存储它们。不能获取寄存器变量的地址。
 * 2. 具有静态存储期的变量可以具有外部链接、内部链接或无链接。在同一个文件所有函数的外部声明的变量是外部变量，
 *    具有文件作用域、外部链接和静态存储期。如果在这种声明前面加上关键字static，那么其声明的变量具有文件作用域、
 *    内部链接和静态存储期。如果在函数中用static声明一个变量，则该变量具有块作用域、无链接、静态存储期。
 * 3. 具有自动存储期的变量，程序在进入该变量的声明所在块时才为其分配内存，在退出该块时释放之前分配的内存。
 *    如果未初始化，自动变量中是垃圾值。程序在编译时为具有静态存储期的变量分配内存，并在程序的运行过程中一直保留这块内存。
 *    如果未初始化，这样的变量会被设置为 0。
 * 4. 具有块作用域的变量是局部的，属于包含该声明的块私有。具有文件作用域的变量对文件（或翻译单元）中位于其声明后面的所有函数可见。
 *    具有外部链接的文件作用域变量，可用于该程序的其他翻译单元。具有内部链接的文件作用域变量，只能用于其声明所在的文件内。
 */
int main(int argc, char *argv[]) {
  int value; // 自动变量
  register int i; // 寄存器变量
  printf("Enter a positive integer(0 to quit): ");
  while (scanf("%d", &value) == 1 && value > 0) {
    ++count; // 使用文件作用域变量
    for (i = value; i >= 0; i--) {
      accumulate(i);
    }
    printf("Enter a positive integer(0 to quit): ");
  }
  report_count();
  return 0;
}

void report_count() {
  printf("Loop executed %d times\n", count);
}