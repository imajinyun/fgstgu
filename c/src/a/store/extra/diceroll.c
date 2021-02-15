#include "diceroll.h"
#include <stdio.h>
#include <stdlib.h>

int roll_count = 0; // 外部链接（源代码文件中的声明，定义式声明）

static int rollem(int side) {
  int roll;
  roll = rand() % side + 1;
  roll_count++;
  return roll;
}

int roll_n_dice(int dice, int side) {
  int d;
  int total = 0;
  if (side < 2) {
    printf("Need at least 2 dice\n");
    return -2;
  }
  if (dice < 1) {
    printf("Need at least 1 die\n");
    return -1;
  }
  for (d = 0; d < dice; d++) { total += rollem(side); }
  return total;
}
