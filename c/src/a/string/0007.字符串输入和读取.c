#include <stdio.h>

#define STRLEN 81


/**
 * 字符串输入和读取。
 *
 * C11 标准委员会从标准中废除了 gets() 函数，
 * 但是编译器为了能兼容以前的代码，大部分都继续支持 gets() 函数。
 *
 * 1. 如果输入的字符串过长，会导致缓冲区溢出（buffer overflow），即多余的字符超出了指定的目标空间；
 * 2. 如果这些多余的字符只是占用了尚未使用的内存，就不会立即出现问题；
 * 3. 如果它们擦写掉程序中的其他数据，会导致程序异常中止；或者还有其他情况；
 */
int main(int argc, char *argv[]) {
  // warning: this program uses gets(), which is unsafe.

  char words[STRLEN];
  puts("Please enter a string:");
  gets(words);
  printf("Your string twice:\n");
  printf("%s\n", words);
  puts(words);
  puts("Done!");
  return 0;
}
