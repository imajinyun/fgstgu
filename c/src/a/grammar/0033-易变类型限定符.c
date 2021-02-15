#include <stdio.h>

void test1() {
  volatile int v1;  // v1 是一个易变的位置
  volatile int *v2; // v2 是一个指向易变的位置的指针
}

/**
 * 智能的（进行优化的）编译器会注意到以上代码使用了两次 x，但并未改变它的值。
 * 于是编译器把 x 的值临时存储在寄存器中，然后在 a 需要使用 x 时，才从寄存器
 * 中（而不是从原始内存位置上）读取 x 的值，以节约时间。这个过程被称为高速缓存（caching）。
 * 通常，高速缓存是个不错的优化方案，但是如果一些其他代理在以上两条语句之间改变了 x 的值，
 * 就不能这样优化了。如果没有 volatile 关键字，编译器就不知道这种事情是否会发生。
 * 因此，为安全起见，编译器不会进行高速缓存。这是在 ANSI 之前的情况。
 * 现在，如果声明中没有 volatile 关键字，编译器会假定变量的值在使用过程中不变，然后再尝试优化代码。
 */
void test2() {
  int x = 99;
  int a, b;
  a = x;
  // 一些不使用 x 的代码
  b = x;
}

void test3() {
  volatile const int v1;
  const volatile int *v2;
}

int main(int argc, char *argv[]) {
  printf("volatile 类型限定符\n");
  test1();
  test2();
  test3();
  return 0;
}
