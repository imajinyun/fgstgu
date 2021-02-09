#include <stdio.h>

#define STRLEN 14

/**
 * 更好的字符串输入和读取。
 *
 * fgets() 和 gets() 的区别如下：
 *
 * 1. fgets() 函数的第 2 个参数指明了读入字符的最大数量。
 *   如果该参数的值是 n，那么 fgets() 将读入 n-1 个字符，或者读到遇到的第一个换行符为止；
 * 2. 如果 fgets() 读到一个换行符，会把它存储在字符串中。这点与 gets() 不同，gets() 会丢弃换行符；
 * 3. fgets() 函数的第 3 个参数指明要读入的文件。如果读入从键盘输入的数据，
 *   则以 stdin（标准输入）作为参数，该标识符定义在 stdio.h 中；
 */
int main(int argc, char *argv[]) {
  char words[STRLEN];
  puts("Please enter a string:");
  fgets(words, STRLEN, stdin);
  printf("Your string twice(puts(), then fputs()):\n");
  puts(words);
  fputs(words, stdout);
  puts("Please enter another string:");
  fgets(words, STRLEN, stdin);
  printf("Your string twice(puts(), then fputs()):\n");
  puts(words);
  fputs(words, stdout);
  puts("Done!");
  return 0;
}
