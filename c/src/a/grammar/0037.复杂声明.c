#include <stdio.h>

void test1() {
  int board[8][8]; // 声明一个内含 int 数组的数组
  int **ptr;       // 声明一个指向指针的指针，被指向的指针指向 int
  int *risks[10];  // 声明一个内含 10 个元素的数组，每个元素都是一个指向 int 的指针
  int *off[3][4];  // 声明一个 3*4 的二维数组，每个元素都是指向 int 的指针
  int(*uuf)[3][4]; // 声明一个指向 3*4 二维数组的指针，该数组中内含 int 类型的值
  int(*uof[3])[4]; // 声明一个内含 3 个指针元素的数组，其中每个指针都指向一个内含 4 个 int 类型元素的数组
  puts("1. 数组名后面的 [] 和函数名后面的 () 具有相同的优先级。它们比 *（解引用运算符）的优先级高；\n"
       "2. [] 和 () 的优先级相同，由于都是从左往右结合；\n");
}

void test2() {
  char *afunc(int);      // 返回字符指针的函数
  char (*bfunc)(int);    // 指向函数的指针，该函数的返回类型为 char
  char (*cfunc[3])(int); // 内含 3 个指针的数组，每个指针都指向返回类型为 char 的函数
}

void test3() {
  typedef int arr[5];
  typedef arr *parr;
  typedef parr qarr[10];
  arr ary; // ary 是一个内含 5 个 int 类型值的数组
  parr p;  // p 是一个指向数组的指针，该数组内含 5 个 int 类型的值
  qarr q;  // q 是一个内含 10 个指针的数组，每个指针都指向一个内含 5 个 int 类型值的数组
}

int main(int argc, char *argv[]) {
  test1();
  test2();
  test3();
  return 0;
}
