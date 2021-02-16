#include <stdio.h>
#include <stdlib.h>

#define CNTL_Z '\032' // DOS 文本文件中的文件结尾标记
#define LENGTH 81

/**
 * 随机访问。
 *
 * -> ANSI C 规定，对于文本模式，ftell() 返回的值可以作为 fseek() 的第 2 个参数。
 *    对于 MS-DOS，ftell() 返回的值把 \r\n 当作一个字节计数。
 *
 * | call | result |
 * |------|--------|
 * | fseek(file, 0L, SEEK_SET) | 定位至文件开始处 |
 * | fseek(file, 0L, SEEK_CUR) | 保持当前位置不动 |
 * | fseek(file, 0L, SEEK_END) | 定位至文件末尾处 |
 * | fseek(file, ftell-pos, SEEK_SET) | 到距文件开始处 ftell-pos 的位置，ftell-pos 是 ftell() 的返回值 |
 */
int main(int argc, char *argv[]) {
  FILE *fp;
  char file[LENGTH];
  char ch;
  long count, last;
  puts("Enter the name of the file to be processed: ");
  scanf("%80s", file);
  if ((fp = fopen(file, "rb")) == NULL) {
    printf("reverse can't open %s\n", file);
    exit(EXIT_FAILURE);
  }
  fseek(fp, 0L, SEEK_END); // 定位到文件末尾，即把当前位置设置为距文件末尾 0 字节偏移量
  last = ftell(fp);        // 把从文件开始处到文件结尾的字节数赋给 last

  // 第 1 轮迭代，把程序定位到文件结尾的第 1 个字符（即，文件的最后一个字符）。
  // 然后，程序打印该字符。下一轮迭代把程序定位到前一个字符，并打印该字符。
  // 重复这一过程直至到达文件的第 1 个字符，并打印。
  for (count = 1L; count <= last; count++) {
    fseek(fp, -count, SEEK_END); // 回退
    ch = getc(fp);
    if (ch != CNTL_Z && ch != '\r') { // MS-DOS 文件
      putchar(ch);
    }
  }
  putchar('\n');
  fclose(fp);
  return 0;
}
