#include <stdio.h>

// 各数据类型所占的字节数。
int main(int argc, char *argv[]) {
  printf("char is %lu bytes\n", sizeof(char)); // 1
  printf("unsigned char is %lu bytes\n", sizeof(unsigned char)); // 1
  printf("short is %lu bytes\n", sizeof(short)); // 2
  printf("unsigned short is %lu bytes\n", sizeof(unsigned short)); // 2
  printf("int is %lu bytes\n", sizeof(int)); // 4
  printf("unsigned int is %lu bytes\n", sizeof(unsigned int)); // 4
  printf("long is %lu bytes\n", sizeof(long)); // 8
  printf("long long is %lu bytes\n", sizeof(long long)); // 8
  printf("float is %lu bytes\n", sizeof(float)); // 4
  printf("double is %lu bytes\n", sizeof(double)); // 8
  printf("long double is %lu bytes\n", sizeof(long double)); // 16
  return 0;
}
