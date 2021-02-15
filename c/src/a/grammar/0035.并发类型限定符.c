#include <stdatomic.h>
#include <stdio.h>

void test1() {
  int i;  // 普通声明
  i = 11; // 普通赋值
}

void test2() {
  // 在 i 中存储 11 是一个原子过程，其他线程不能访问 i
  _Atomic int i;
  atomic_store(&i, 11); // stdatomic.h 中的宏
}

int main(int argc, char *argv[]) {
  printf("当一个线程对一个原子类型的对象执行原子操作时，其他线程不能访问该对象。\n");
  test1();
  test2();
  return 0;
}
