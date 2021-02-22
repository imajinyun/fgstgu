#include <stdio.h>

inline static void eatline() {
  while (getchar() != '\n') { continue; }
}

int main(int argc, char *argv[]) {
  puts("编译器优化内联函数必须知道该函数定义的内容。这意味着内联函数定义与函数调用必须在同一个文件中。");
  printf("Please input enter");
  eatline();
  return 0; }
