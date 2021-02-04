#include <stdio.h>

int main(int argc, char *argv[]) {
  printf("请输入一个数字，输出的数字表示星期几：");

  int i;
  scanf("%d", &i);

  switch (i) {
    case 1:
      printf("输入的 %d 表示 Monday\n", i);
      break;
    case 2:
      printf("输入的 %d 表示 Tuesday\n", i);
      break;
    case 3:
      printf("输入的 %d 表示 Wednesday\n", i);
      break;
    case 4:
      printf("输入的 %d 表示 Thursday\n", i);
      break;
    case 5:
      printf("输入的 %d 表示 Friday\n", i);
      break;
    case 6:
      printf("输入的 %d 表示 Saturday\n", i);
      break;
    case 7:
      printf("输入的 %d 表示 Sunday\n", i);
      break;
    default:
      printf("请输入 1 到 7 之间的数字\n");
  }
  return 0;
}
