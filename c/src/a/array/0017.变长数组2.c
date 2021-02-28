#include <stdio.h>

int main(int argc, char *argv[]) {
  int num;

  printf("Please enter a number: ");
  scanf("%d", &num);
  scanf("%*[^\n]");
  scanf("%*c");

  char ch[num];
  printf("Please enter  a string: ");
  fgets(ch, num + 1, stdin);
  puts(ch);
  return 0;
}
