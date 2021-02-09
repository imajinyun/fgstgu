#include <stdio.h>
#include <string.h>

#define MSG "I'm special"

/**
 * 数组和指针：
 *
 * 1. 初始化数组把静态存储区的字符串拷贝到数组中；
 * 2. 初始化指针只把字符串的地址拷贝给指针；
 */
int main(int argc, char *argv[]) {
  // 静态数据使用的内存与 arr 使用的动态内存不同。不仅值不同，特定编译器甚至使用不同的位数表示两种内存。
  char arr[] = MSG;
  const char *ptr = MSG;
  printf("\naddress of \"I'm special\": %p\n", "I'm special");
  printf("             address arr: %p\n", arr);
  printf("             address ptr: %p\n", ptr);
  printf("          address of MSG: %p\n", MSG);
  printf("address of \"I'm special\": %p\n", "I'm special");
  return 0;
}
