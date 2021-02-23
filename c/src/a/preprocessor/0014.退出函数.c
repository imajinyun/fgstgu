#include <stdio.h>
#include <stdlib.h>

void sign_off();
void too_bad();

int main(int argc, char *argv[]) {
  int n;
  atexit(sign_off); // 注册 sign_off 函数
  printf("Enter an integer: ");
  if (scanf("%d", &n) != 1) {
    puts("That's no integer!");
    atexit(too_bad); // 注册 too_bad 函数
    exit(EXIT_FAILURE);
  }
  printf("%d is %s.\n", n, (n % 2 == 0) ? "even" : "odd");
  return 0;
}

void sign_off() {
  puts("Thus terminates another magnificent program from");
  puts("SeeSaw software!");
}

void too_bad() {
  puts("SeeSaw software extends its heartfelt condolences");
  puts("to you upon the failure of your program.");
}
