#include <stdio.h>
#include <unistd.h>

/**
 * 清空输出缓冲区。
 **/
int main(int argc, char *argv[]) {
  printf("清空输出缓冲区1");
  fflush(stdout);
  sleep(2);
  printf("\n🎉 Hello World\n");
  return 0;
}
