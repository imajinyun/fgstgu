#include <stdio.h>
#include <string.h>

/**
 * 数组和指针的区别。
 */
int main(int argc, char *argv[]) {
  char a[] = "I love Tillie!";
  const char *b = "I love Millie!";
  int alen = strlen(a) - 8, blen = strlen(b) - 8;

  // 两者都可以使用数组表示法
  for (int i = 0; i < alen; i++) { putchar(a[i]); }
  putchar('\n');
  for (int i = 0; i < blen; i++) { putchar(b[i]); }
  putchar('\n');

  // 两者都能进行指针加法操作
  for (int i = 0; i < alen; i++) { putchar(*(a + i)); }
  putchar('\n');
  for (int i = 0; i < blen; i++) { putchar(*(b + i)); }
  printf("\n\n");

  // 只有指针表示法可以进行递增操作
  while (*(b) != '\0') { putchar(*(b++)); } // I love Millie!
  putchar('\n');

  b = a;
  while (*(b) != '\0') { putchar(*(b++)); } // I love Tillie!
  putchar('\n');

  a[7] = 'P';
  *(a + 8) = 'y';
  for (int i = 0; i < strlen(a); i++) { putchar(*(a + i)); } // I love Pyllie!

  // 编译器可能允许这样做，但是对当前的 C 标准而言，这样的行为是未定义的，这样的语句可能导致内存访问错误。
  while (*(b) != '\0') { putchar(*(b++)); } // 没有任何输出
  printf("\n");

  // 如果打算修改字符串，就不要用指针指向字符串字面量。
  // 把非 const 数组初始化为字符串字面量却不会导致类似的问题，因为数组获得的是原始字符串的副本。
  const char *p = "Klingon"; // 推荐用法
  return 0;
}
