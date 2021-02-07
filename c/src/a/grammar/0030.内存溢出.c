#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
  printf("内存溢出的实例（运行后要记得关闭掉）：\n");
  while (1) { malloc(1048 * 8); }
  return 0;
}
