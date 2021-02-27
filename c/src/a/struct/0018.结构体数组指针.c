#include <stdio.h>

struct stu {
  char *name;
  int num;
  int age;
  char group;
  float score;
} stus[] = {{"ChenJi", 1002, 19, 'C', 560},
            {"LiZhen", 1010, 19, 'A', 601.5},
            {"ZhaoPi", 1011, 18, 'A', 620.5},
            {"HaoTao", 1021, 19, 'B', 590.5},
            {"WanYun", 1037, 18, 'D', 605.5}};
struct stu *ps;

int main(int argc, char *argv[]) {
  int n = sizeof(stus) / sizeof(struct stu);
  printf("Name     Num   Age   Group Score\n");
  for (ps = stus; ps < stus + n; ps++) {
    char *format = "%-9s%-6d%-6d%-2c%9.1f\n";
    printf(format, ps->name, ps->num, ps->age, ps->group, ps->score);
  }
  return 0;
}
