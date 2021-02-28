#include <stdio.h>

int main(int argc, char *argv[]) {
  // : 后面的数字用来限定成员变量占用的位数
  struct bs {
    unsigned m; // 成员 m 没有限制，根据数据类型即可推算出它占用 4 个字节（Byte）的内存
    unsigned n : 4;
    unsigned char c : 6;
  } a = {0xad, 0xE, '$'};

  // n、c 的取值范围非常有限，数据稍微大些就会发生溢出，所以要注意取值范围
  a.m = 0xb8901c;
  a.n = 0x00000f;
  a.c = '?';

  printf("sizeof(struct bs) = %lu\n", sizeof(struct bs));
  printf("a.m = %#x, a.n = %#x, a.c = %#x\n", a.m, a.n, a.c);
  printf("\n");

  puts("C 语言标准规定，位域的宽度不能超过它所依附的数据类型的长度。\n"
       "通俗地讲，成员变量都是有类型的，这个类型限制了成员变量的最大长度，: 后面的数字不能超过这个长度。\n"
       "位域技术就是在成员变量所占用的内存中选出一部分位宽来存储数据。\n");

  puts("C 语言标准还规定，只有有限的几种数据类型可以用于位域。\n"
       "在 ANSI C 中，这几种数据类型是 int、signed int 和 unsigned int（int 默认就是 signed int）。\n"
       "到了 C99，_Bool 也被支持了。");
  return 0;
}
