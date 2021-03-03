#include "diceroll.h"
#include <stdio.h>
#include <stdlib.h>

int roll_count = 0; // 外部链接（源代码文件中的声明，定义式声明，该变量具有文件作用域）

static int rollem(int side) {
  int roll;
  roll = rand() % side + 1;
  roll_count++;
  return roll;
}

int roll_n_dice(int dice, int side) {
  int total = 0;
  if (side < 2) {
    printf("Need at least 2 dice\n");
    return -2;
  }
  if (dice < 1) {
    printf("Need at least 1 die\n");
    return -1;
  }
  for (int i = 0; i < dice; i++) { total += rollem(side); }
  return total;
}
