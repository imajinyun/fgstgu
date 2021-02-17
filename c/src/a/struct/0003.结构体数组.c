#include <stdio.h>

struct {
  char *name;
  int num;
  int age;
  char group;
  float score;
} stu[] = {{"Zhang San", 1002, 19, 'C', 560},
           {"Li Si", 1010, 19, 'A', 601.5},
           {"Chen Zhen", 1011, 18, 'A', 620.5},
           {"Hao Tao", 1021, 19, 'B', 590.5},
           {"Wang Yong", 1037, 18, 'D', 605.5}};

int main(int argc, char *argv[]) {
  int i, num = 0;
  float sum = 0;

  // sizeof(stu) / sizeof(stu[0]) 求出的是数组的总长度，而不是数组中存放的有意义的数据的个数。
  unsigned long len = sizeof(stu) / sizeof(stu[0]);

  for (i = 0; i < len; i++) {
    sum += stu[i].score;
    if (stu[i].score < 600) { num++; }
  }
  printf("sum = %.2f\n", sum);
  printf("avg = %.2f\n", sum / len);
  printf("num = %d\n", num);
  return 0;
}
