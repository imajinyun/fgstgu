#include <stdio.h>

#define LIMIT 100           // LIMIT 是已定义的
#define GOOD                // GOOD 是已定义的
#define A(x) ((-(x)) * (x)) // A 是已定义的
int x;                      // x 不是宏，属于未定义的
#undef GOOD                 // GOOD 取消定义，是未定义的

#ifdef ISSIGN
#define VERSION 11 // 如果已使用 #define 定义了 ISSIGN，则执行此指令
#else
#define VERSION 10 // 如果未使用 #define 定义过 ISSIGN，则执行此指令
#endif

int main(int argc, char *argv[]) {
  puts("注意：\n"
       "  #define 宏的作用域从它在文件中的声明处开始，直到用 #undef 指令取消宏为止，\n"
       "  或延伸至文件尾（以二者中先满足的条件作为宏作用域的结束）。另外还要注意，如果宏通过头文件引\n"
       "  入，那么 #define 在文件中的位置取决于 #include 指令的位置。");
  return 0;
}
