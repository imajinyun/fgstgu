#include <stdio.h>
#include <unistd.h>

int main(int argc, char *argv[]) {
  char ch1;
  printf("请输入一个字符：");
  ch1 = getchar();
  printf("ch1 = %c\n", ch1);

  printf("🎉 Hello");
  sleep(2);
  puts(" World!\n");
  puts("缓冲区的刷新遵循以下的规则：\n"
       "  1. 不管是行缓冲还是全缓冲，缓冲区满时会自动刷新；\n"
       "  2. 行缓冲遇到换行符 \\n 时会刷新；\n"
       "  3. 关闭文件时会刷新缓冲区；\n"
       "  4. 程序关闭时一般也会刷新缓冲区，这个是由标准库来保障的；\n"
       "  5. 使用特定的函数也可以手动刷新缓冲区；");
  return 0;
}
