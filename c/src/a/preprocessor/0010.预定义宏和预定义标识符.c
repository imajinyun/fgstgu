#include <stdio.h>

#pragma message "Only show message at compileing..."

void test1();
void test2();

int main(int argc, char *argv[]) {
  printf("The file is %s.\n", __FILE__);
  printf("The data is %s.\n", __DATE__);
  printf("The time is %s.\n", __TIME__);
  printf("The version is %ld.\n", __STDC_VERSION__);
  printf("This is line %d.\n", __LINE__);
  printf("This function is %s.\n", __func__);

  test1();
  test2();
  return 0;
}

void test1() {
  printf("This function is %s.\n", __func__);
  printf("This is line %d.\n", __LINE__);
}

void test2() {
#if __STDC_VERSION__ != 201112L
#error Not C11
#endif
}
