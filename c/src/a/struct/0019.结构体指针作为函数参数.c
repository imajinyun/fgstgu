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

void avg(struct stu *ps, int n);

int main(int argc, char *argv[]) {
  int n = sizeof(stus) / sizeof(struct stu);
  avg(stus, n);
  return 0;
}

void avg(struct stu *ps, int n) {
  int i, num;
  float avg, sum = 0;
  for (i = 0; i < n; i++) {
    sum += (ps + i)->score;

    if ((ps + i)->score >= 600) { num++; }
  }
  printf("班级总成绩为：%.2f\n", sum);
  printf("班级平均分为：%.2f\n", sum / n);
  printf("600 分以上人数为：%d\n", num);
}
