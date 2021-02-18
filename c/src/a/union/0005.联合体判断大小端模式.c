#include <stdio.h>

int main(int argc, char *argv[]) {
  union {
    int i;
    char c;
  } data;
  data.i = 0x00000001;

  if (data.c == 1) {
    printf("Little-endian\n");
  } else {
    printf("Big-endian\n");
  }
  printf("\n");

  puts("大小端概念：\n"
       "1. 大端模式：是指数据的高字节保存在内存的低地址中，而数据的低字节保存在内存的高地址中；\n"
       "2. 小端模式，是指数据的高字节保存在内存的高地址中，而数据的低字节保存在内存的低地址中；\n");

  puts("联合体的各个成员是共用一段内存的：\n"
       "1. 如果 1 是数据的低位，那么 1 被存储在 data 的低字节，就是小端模式，这个时候 data.c 的值就是 1；\n"
       "2. 如果 1 是数据的高位，那么 1 被存储在 data 的高字节，就是大端模式，这个时候 data.c 的值就是 0；");
  return 0;
}
