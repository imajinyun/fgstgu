#include <stdio.h>

void ToUpper(char *);
void ToLower(char *);
int rounder(double);
void (*pfunc)(char *);

int main(int argc, char *argv[]) {
  printf("函数指针使用声明：\n");
  pfunc = ToUpper; // 有效，ToUpper 是该类型函数的地址
  pfunc = ToLower; // 有效，ToUpper 是该类型函数的地址
  // pfunc = rounder;   // 无效，rounder 与指针类型不匹配
  // pfunc = ToLower(); // 无效，ToLower() 不是地址
  return 0;
}

void ToUpper(char *s) {}
void ToLower(char *s) {}
int rounder(double x) { return 0; }
