#include <stdio.h>

void test1() {
  float f = 3.14;

  // 创建的指针 p1 指向不能被改变的值，而 p1 本身的值可以改变。
  const float *p1; // p1 是一个 float 类型的 const 值
  float const *p2; // p2 是一个 float 类型的 const 值

  // *p1 = f; // error: read-only variable is not assignable
  p1 = &f; // no problem

  // *p2 = f; // error: read-only variable is not assignable
  p2 = &f;
  printf("p1 addr is %p\n", p1);
  printf("p2 addr is %p\n", p2);
}

void test2() {
  float f = 3.1415926;

  // 创建的指针 p3 本身的值不能更改。p3 必须指向同一个地址，但是它所指向的值可以改变。
  float *const p3; // p3 是一个 const 指针

  *p3 = f;

  // error: cannot assign to variable 'p3' with const-qualified type 'float *const'
  // note: variable 'p3' declared const here
  // p3 = &f;
  printf("p3 addr is %p\n", p3);
}

void test3() {
  float f = 3.141592653589793;

  // 创建的指针 p4 既不能指向别处，它所指向的值也不能改变。
  const float *const p4;
  // *p4 = f; // error: read-only variable is not assignable

  // error: cannot assign to variable 'p4' with const-qualified type 'const float *const'
  // note: variable 'p4' declared const here
  // p4 = &f;
  printf("p4 addr is %p\n", p4);
}

/**
 * 指针中的常量限定符。
 *
 * -> const 放在 * 左侧任意位置，限定了指针指向的数据不能改变。
 * -> const 放在 * 的右侧，限定了指针本身不能改变。
 */
int main(int argc, char *argv[]) {
  test1();
  test2();
  test3();
  return 0;
}
