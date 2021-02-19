#include <stdio.h>

void ToUpper(char *);
void ToLower(char *);
void (*pfunc)(char *);

int main(int argc, char *argv[]) {
  printf("函数指针声明：\n");
  char mis[] = "Hello World";
  pfunc = ToUpper; // 有效，ToUpper 是该类型函数的地址
  (*pfunc)(mis);   // 把 ToUpper 作用于 mis（语法一）
  pfunc = ToLower; // 有效，ToUpper 是该类型函数的地址
  pfunc(mis);      // 把 ToLower 作用于 mis（语法二）
  return 0;
}

void ToUpper(char *s) {}
void ToLower(char *s) {}
