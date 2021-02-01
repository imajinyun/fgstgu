#include <stdio.h>

void flush() {
  scanf("%*[^\n]");
  scanf("%*c");
}

/**
 * 匹配特定的字符。
 *
 * -> abcXYZ@#87edf
 * -> c c++ java python go javascript
 * -> http://baidu.com http://tmall.com
 */
int main(int argc, char *argv[]) {
  char s1[30], s2[30], s3[30];
  printf("请输入匹配的字符串：\n");
  scanf("%[^0-9]", s1);
  flush();

  // scanf() 也能读取带空格的字符串
  scanf("%[^\n]", s2);
  flush();

  scanf("%30[^0-9\n]", s3);
  printf("s1 = %s\ns2 = %s\ns3 = %s\n", s1, s2, s3);
  return 0;
}
