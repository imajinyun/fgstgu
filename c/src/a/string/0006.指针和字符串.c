#include <stdio.h>

int main(int argc, char *argv[]) {
  // 如果编译器不识别 %p，用 %u 或 %lu 代替 %p。
  const char *a = "Don't be a fool!";
  const char *b;

  // 语句 b = a; 把 a 的值赋给 b，即让 b 也指向 a 指向的字符串。
  b = a;

  // 所谓指针的值就是它存储的地址。
  printf("b = %s\n", b);
  printf("a = %s, &a = %p, a pointer value = %p\n", a, &a, a);
  printf("b = %s, &b = %p, b pointer value = %p\n", b, &b, b);
  return 0;
}
