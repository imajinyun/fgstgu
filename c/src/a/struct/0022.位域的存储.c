#include <stdio.h>

void test1();
void test2();
void test3();

int main(int argc, char *argv[]) {
  test1();
  test2();
  test3();
  return 0;
}

void test1() {
  // 当相邻成员的类型相同时，如果它们的位宽之和小于类型的 sizeof 大小，那么后面的成员紧邻前一个成员存储，直到不能容纳为止；
  // 如果它们的位宽之和大于类型的 sizeof 大小，那么后面的成员将从新的存储单元开始，其偏移量为类型大小的整数倍。
  struct bs {
    unsigned x : 6;
    unsigned y : 12;
    unsigned z : 4;
  };
  // x、y、z 的类型都是 unsigned int，sizeof 的结果为 4 个字节（Byte），也即 32 个位（Bit）
  printf("sizeof(struct bs) = %lu\n", sizeof(struct bs));
  puts("x、y、z 的位宽之和为 6+12+4 = 22，小于 32，所以它们会挨着存储，中间没有缝隙。\n");
  printf("\n");
}

void test2() {
  // 当相邻成员的类型不同时，不同的编译器有不同的实现方案，GCC 会压缩存储，而 VC/VS 不会。
  struct bs {
    unsigned x : 12;
    unsigned char y : 4;
    unsigned z : 4;
  };
  printf("sizeof(struct bs) = %lu\n", sizeof(struct bs));

  printf("\n");
}

void test3() {
  // 如果成员之间穿插着非位域成员，那么不会进行压缩。
  struct bs {
    unsigned x : 12;
    unsigned y;
    unsigned z : 4;
  };
  printf("sizeof(struct bs) = %lu\n", sizeof(struct bs));
  puts("位域成员往往不占用完整的字节，有时候也不处于字节的开头位置，\n"
       "因此使用 & 获取位域成员的地址是没有意义的，C 语言也禁止这样做。\n"
       "地址是字节（Byte）的编号，而不是位（Bit）的编号。");
}
