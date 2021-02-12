#include <stdio.h>

#define DEF "I am a #define string"

void test1() {
  char s1[80] = "An array was initialized to me";
  const char *s2 = "A pointer was initialized to me";
  puts("I'm a argument to puts()");
  puts(DEF);
  puts(s1);
  puts(s2);
  puts(&s1[3]);
  puts(s2 + 2);
}

void test2() {
  // 由于 b 缺少一个表示结束的空字符，所以它不是一个字符串，
  // 因此 puts() 不知道在何处停止。它会一直打印 b 后面内存中的内容，
  // 直到发现一个空字符为止。为了让 puts() 能尽快读到空字符。
  char a[] = "Side A";
  char b[] = {'W', 'O', 'W', '!'}; // 不要这样做
  char c[] = "Side B";

  // 编译器把 a 数组存储在 b 数组之后，所以 puts() 一直输出至遇到 a 中的空字符，
  // 所以使用的编译器不同输出的内容可能不同。
  // -> WOW!Side A
  puts(b); // b 不是一个字符串
}

void test3() {
  char *s = "Hello World";
  printf("*s = %s, DEF = %s\n", s, DEF);
}

void test4() {
  char line[81];
  while (fgets(line, 81, stdin)) { fputs(line, stdout); }
}

int main(int argc, char *argv[]) {
  test1();
  printf("\n");
  test2();
  printf("\n");
  test3();
  printf("\n");
  test4();
  return 0;
}
