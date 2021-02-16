#include <stdio.h>

int main(int argc, char *argv[]) {
  struct {
    char *name;  // 姓名
    int num;     // 学号
    int age;     // 年龄
    char group;  // 属组
    float score; // 成绩
  } stu;

  stu.name = "Jack";
  stu.num = 1201;
  stu.age = 28;
  stu.group = 'A';
  stu.score = 580.5;

  printf("姓名：%s\n", stu.name);
  printf("学号：%d\n", stu.num);
  printf("年龄：%d\n", stu.age);
  printf("属组：%c\n", stu.group);
  printf("成绩：%.1f\n", stu.score);
  return 0;
}
