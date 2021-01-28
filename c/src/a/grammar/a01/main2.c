#include <stdio.h>

// 输入输出。
int main(int argc, char *argv[]) {
  char name[30];
  int age;
  float height;

  printf("请输入你的姓名年龄身高（多个值请用空格分隔）：");
  scanf("%s %d %f", name, &age, &height);
  printf("姓名：%s 年龄：%d 身高：%.2f\n", name, age, height);
  printf("地址：&name = %p &age = %p &height = %p\n", &name, &age, &height);
  return 0;
}
