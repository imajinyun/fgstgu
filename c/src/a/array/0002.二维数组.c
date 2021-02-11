#include <stdio.h>

int main(int argc, char *argv[]) {
  int a[3][3] = {
    {1},
    {2},
    {3},
  };
  puts("二维数组初始化：");
  for (int p = 0; p < 3; p++) {
    for (int q = 0; q < 3; q++) { printf("%d ", a[p][q]); }
    printf("\n");
  }

  int i, j;
  float sum = 0.0;
  float avg = 0.0;
  float averages[3];
  float scores[5][3] = {
    {70.5, 86, 67.5}, {77, 66.5, 71.5}, {92.5, 83.5, 85}, {78, 64.5, 75.5}, {82.5, 72.5, 89},
  }; // 5 个同学 3 个科目（假定为语数外）

  for (i = 0; i < 3; i++) {
    for (j = 0; j < 5; j++) {
      sum += scores[j][i]; // 计算当前科目的总成绩
    }
    averages[i] = sum / 5; // 当前科目的平均分
    sum = 0.0;
  }
  avg = (averages[0] + averages[1] + averages[2]) / 3;

  puts("\n计算成绩平均数：");
  printf("语文平均分数：%.1f\n", averages[0]);
  printf("数学平均分数：%.1f\n", averages[1]);
  printf("外语平均分数：%.1f\n", averages[2]);
  printf("三科平均分数：%.1f\n", avg);
  return 0;
}
