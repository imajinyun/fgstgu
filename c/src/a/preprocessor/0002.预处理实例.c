#include <stdio.h>

#define TWO 2 /* 可以使用注释 */
#define OW                                                                                                             \
  "Consistency is the last refuge of the unimagina\
tive. - Oscar Wille" /* 反斜杠把该定义延续到下一行 */
#define FOUR TWO *TWO
#define PX printf("x is %d.\n", x)
#define FMT "x is %d.\n"

/**
 * 预处理实例。
 *
 * 预处理器指令从 # 开始运行，到后面的第 1 个换行符为止。
 * 也就是说，指令的长度仅限于一行。
 * 然而，在预处理开始前，编译器会把多行物理行处理为一行逻辑行。
 */
int main(int argc, char *argv[]) {
  int x = TWO;
  PX;
  x = FOUR;
  printf(FMT, x);
  printf("%s\n", OW);
  printf("TWO: OW\n");
  return 0;
}
