#include <stdio.h>
#include <string.h>

int main(int argc, char *argv[]) {
  char ch1[] = "this is a test string1";
  printf("字符数组输出：\n");
  printf("%s\n", ch1);

  int l1 = strlen(ch1);
  ch1[l1] = '@'; // 正确
  printf("%s\n", ch1);
  printf("\n");

  char *ch2 = "this is a test string2";
  printf("字符串常量输出：\n");
  printf("%s\n", ch2);

  ch2 = "Hello World!"; // 正确
  printf("%s\n", ch2);

  int l2 = strlen(ch2);
  // ch2[l2] = '@'; // 错误

  puts("\n字符数组与字符串常量的区别：\n"
       "1. 字符数组存储在全局数据区或栈区，字符数组在定义后可以读取和修改每个字符；\n"
       "2. 字符串常量存储在常量区，只能读取不能修改，任何对它的赋值都是错误的；\n"
       "3. 获取用户输入的字符串就是一个典型的写入操作，只能使用字符数组，不能使用字符串常量；");
  return 0;
}
