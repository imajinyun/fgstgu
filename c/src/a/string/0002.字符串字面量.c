#include <stdio.h>

/**
 * 字符串字面量：
 *
 * 1. 用双引号括起来的内容称为字符串字面量（string literal），也叫作字符串常量（string constant）；
 * 2. 双引号中的字符和编译器自动加入末尾的 \0 字符，都作为字符串存储在内存中；
 */
int main(int argc, char *argv[]) {
  char a[50] = "🎉 Hello World, "
               "Hello C!";
  char b[50] = "🎉 Hello World, Hello C!";
  puts(a);
  puts(b);
  printf("%s, %p, %c\n", "🎉 Hello", " World, ", *"Hello C!");
  return 0;
}
