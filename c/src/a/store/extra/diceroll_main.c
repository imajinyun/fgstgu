#include "diceroll.h"
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int main(int argc, char *argv[]) {
  int dice, roll;
  int side, status;
  srandom((unsigned int) time(0)); // 随机种子
  printf("Enter the number of sides per die (0 to stop)\n");
  while (scanf("%d", &side) == 1 && side > 0) {
    printf("How many dice?\n");
    if ((status = scanf("%d", &dice)) != 1) {
      if (status == EOF) {
        break;
      } else {
        printf("You should have entered an integer");
        printf("Let's begin again\n");
        while (getchar() != '\n') { continue; }
        printf("How many sides (0 to stop)?\n");
        continue;
      }
    }
    roll = roll_n_dice(dice, side);
    printf("You have rolled a %d using %d %d-sided dice\n", roll, dice, side);
    printf("How many sides (0 to stop)?\n");
  }
  printf("The rollem() function was called %d times\n", roll_count); // 使用外部变量
  printf("Good fortune to your!\n");
  return 0;
}
