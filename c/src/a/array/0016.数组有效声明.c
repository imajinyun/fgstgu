#include <stdio.h>

// int sum1(int arr[][], int rows);   // 错误的声明
int sum2(int arr[][4], int rows);  // 有效的声明
int sum3(int arr[3][4], int rows); // 有效的声明，3 将被忽略

typedef int arr[4]; // arr 是一个内含 4 个 int 的数组
typedef arr ary[3]; // ary 是一个内含 3 个 arr 的数组
int sum10(ary arr, int rows);
int sum11(int arr[3][4], int rows);
int sum12(int arr[][4], int rows); // 标准形式

// 以下两种写法都是等价的
// arr 指向一个 2*3*4 的 int 数组
int sum4d(int arr[][2][3][4], int rows);
int sum4d(int (*arr)[2][3][4], int rows);

int main(int argc, char *argv[]) {
  printf("数组有效声明\n");
  printf("一般而言，声明一个指向 N 维数组的指针时，只能省略最左边方括号中的值\n");
  printf("因为第 1 对方括号只用于表明这是一个指针，而其他的方括号则用于描述指针所指向数据对象的类型\n");
  return 0;
}
