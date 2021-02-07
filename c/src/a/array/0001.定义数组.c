#include <stdio.h>
#include <stdlib.h>

#define SIZE 10

#define random(min, max) (min + (rand() % (max - min + 1)))

int main(int argc, char *argv[]) {
  char strings[SIZE] = {0};
  int numbers[SIZE] = {0};
  float scores[SIZE] = {0};

  puts("字符数组：");
  for (int i = 0; i < SIZE; i++) {
    strings[i] = i + 'a';
    printf("%c ", strings[i]);
  }

  puts("\n\n整型数组：");
  for (int j = 0; j < SIZE; j++) {
    numbers[j] = j;
    printf("%d ", numbers[j]);
  }

  puts("\n\n浮点数组：");
  for (int k = 0; k < SIZE; k++) {
    scores[k] = (float) k * random(1, 10);
    printf("%.2f ", scores[k]);
  }

  int a[5] = {1, 2, 3, 4, 5};
  int b[5] = {2672, 1283, 683, 99, 890};
  int c[5] = {22, 338, 2989, 9989, 238};
  int d[5] = {0, 12, 3389, 238, 23};
  int e[5] = {9892, 233, 23, 7, 1234};

  puts("\n\n数组矩阵：");
  printf("%-9d %-9d %-9d %-9d %-9d\n", a[0], a[1], a[2], a[3], a[4]);
  printf("%-9d %-9d %-9d %-9d %-9d\n", b[0], b[1], b[2], b[3], b[4]);
  printf("%-9d %-9d %-9d %-9d %-9d\n", c[0], c[1], c[2], c[3], c[4]);
  printf("%-9d %-9d %-9d %-9d %-9d\n", d[0], d[1], d[2], d[3], d[4]);
  printf("%-9d %-9d %-9d %-9d %-9d\n", e[0], e[1], e[2], e[3], e[4]);

  printf("\n");
  puts("传统数组的缺点：\n"
       "1. 数据长度必须事先制定，且只能为常整数，不能是变量\n"
       "2. 传统形式的数组，该数组的内存程序员无法手动释放。在一个函数运行期间，"
       "系统为该函数中数组分配的空间会一直存在，直到该函数运行完毕时，数组的空间才会被系统释放\n"
       "3. 数组的长度一旦定义，其长度就不能在更改，数组的长度不能在函数运行的过程中动态的扩充或缩小\n"
       "4. A 函数定义的数组，在 A 函数运行期间可以被其它函数使用，但 A 函数运行完毕后，A "
       "函数中的数组将无法在被其它函数使用，"
       "即不能跨函数使用");

  printf("\n");
  return 0;
}
