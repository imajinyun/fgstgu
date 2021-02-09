#include <stdio.h>

#define STRLEN 10

/**
 * 更好的字符串输入和读取。
 *
 * 删除存储在字符串中的换行符，如果没有换行符，则丢弃数组装不下的字符。
 */
int main(int argc, char *argv[]) {
  char words[STRLEN];
  int i;
  puts("Enter strings(empty line to quit):");
  while (fgets(words, STRLEN, stdin) != NULL && words[0] != '\n') {
    i = 0;
    while (words[i] != '\n' && words[i] != '\0') { i++; }
    if (words[i] == '\n') {
      words[i] = '\0';
    } else {
      while (getchar() != '\n') { continue; }
    }
    puts(words);
  }
  puts("Done!");
  return 0;
}
