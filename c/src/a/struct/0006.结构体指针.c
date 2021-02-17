#include <stdio.h>

int main(int argc, char *argv[]) {
  struct {
    char *name;
    int num;
    int age;
    char group;
    float score;
  } stu = {"Jack", 1044, 19, 'B', 611.5}, *ps = &stu;


  char *format = "姓名：%s，学号：%d，年龄：%d，分组：%c，成绩：%.2f\n";
  printf(format, (*ps).name, (*ps).num, (*ps).age, (*ps).group, (*ps).score);
  printf(format, ps->name, ps->num, ps->age, ps->group, ps->score);
  return 0;
}
