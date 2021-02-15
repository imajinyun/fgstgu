#include <stdio.h>
#include <stdlib.h>

/**
 * 指针 restar 是访问由 malloc() 所分配内存的唯一且初始的方式。
 * 因此，可以用 restrict 关键字限定它。而指针 ptr 既不是访问 arr 数组中数据的初始方式，
 * 也不是唯一方式。所以不用把它设置为 restrict。
 */
void test1() {
  int arr[10];
  int *restrict restar = (int *) malloc(10 * sizeof(int));
  int *ptr = arr;
}

/**
 * 由于之前声明了 restar 是访问它所指向的数据块的唯一且初始的方式，
 * 编译器可以把涉及 restar 的两条语句替换成下面这条语句：
 * restar[n] += 8; // 可以替换
 *
 * 如果把与 ptr 相关的两条语句替换成下面的语句，将导致计算错误：
 * ptr[n] += 8; // 抛出错误，这是因为 for 循环在 ptr 两次访问相同的数据之间，用 arr 改变了该数据的值
 */
void test2() {
  int arr[10];
  int *restrict restar = (int *) malloc(10 * sizeof(int));
  int *ptr = arr;

  for (int n = 0; n < 10; n++) {
    ptr[n] += 5;
    restar[n] += 5;
    arr[n] *= 2;
    ptr[n] += 3;
    restar[n] += 3;
  }

  printf("arr output:\n");
  for (int i = 0; i < 10; i++) { printf("arr[%d] = %d\n", i, arr[i]); }
  printf("\n*restar output:\n");
  for (int j = 0; j < 10; j++) { printf("*restar[%d] = %d\n", j, *(restar++)); }
}

/**
 * 限制类型限定符。
 *
 * -> restrict 限定符还可用于函数形参中的指针。这意味着编译器可以假定
 *    在函数体内其他标识符不会修改该指针指向的数据，而且编译器可以尝试
 *    对其优化，使其不做别的用途。
 *
 * -> restrict 关键字有两个读者：
 *    一个是编译器，该关键字告知编译器可以自由假定一些优化方案。
 *    另一个读者是用户，该关键字告知用户要使用满足 restrict 要求的参数。
 *    总而言之，编译器不会检查用户是否遵循这一限制，但是无视它后果自负。
 */
int main(int argc, char *argv[]) {
  test1();
  test2();
  return 0;
}
