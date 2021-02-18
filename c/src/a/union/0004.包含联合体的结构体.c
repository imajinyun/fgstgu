#include <stdio.h>
#include <stdlib.h>

#define TOTAL 2

struct {
  char name[20];
  int num;
  char sex;
  char profession;
  union {
    float score;
    char course[20];
  } sc;
} info[TOTAL];

/**
 * 包含联合体的结构体。
 *
 * [in]: 张三 1001 A s 89.90
 * [in]: 李四 1002 B t 数学
 */
int main(int argc, char *argv[]) {
  int i;
  for (i = 0; i < TOTAL; i++) {
    printf("输入信息：");
    scanf("%s %d %c %c", info[i].name, &info[i].num, &info[i].sex, &info[i].profession);

    if (info[i].profession == 's') {
      scanf("%f", &info[i].sc.score);
    } else {
      scanf("%s", info[i].sc.course);
    }
    fflush(stdin);
  }
  printf("\n输出信息：\n");

  for (i = 0; i < TOTAL; i++) {
    if (info[i].profession == 's') {
      printf("%-10s%-10d%-10c%-10c%-10.2f\n", info[i].name, info[i].num, info[i].sex, info[i].profession,
             info[i].sc.score);
    } else {
      printf("%-10s%-10d%-10c%-10c%-10s\n", info[i].name, info[i].num, info[i].sex, info[i].profession,
             info[i].sc.course);
    }
  }
  return 0;
}
