#include <stdio.h>

int units = 0; // 外部变量

void critic();

/**
 * 使用局部静态变量。
 *
 * -> 定义：即定义式声明（defining declaration）要为变量分配内存空间；
 * -> 声明：即引用式声明（referencing declaration）不需要为变量分配内存空间；
 *
 * -> 关键字 extern 表明该声明不是定义，因为它指示编译器去别处查询其定义。
 * -> 不要用关键字 extern 创建外部定义，只用它来引用现有的外部定义。
 *
 * 1. units 具有文件作用域、外部链接和静态存储；
 * 2. 声明主要是为了指出该函数要使用这个外部变量；
 * 3. 存储类别说明符 extern 告诉编译器，该函数中任何使用 units 的地方都引用同一个定义在函数外部的变量；
 * 4. 注意：main() 和 critic() 使用的都是外部定义的 units；
 */
int main(int argc, char *argv[]) {
  extern int units; // 可选的重复声明
  printf("How many pounds to a firkin of butter?\n");
  scanf("%d", &units);
  while (units != 50) { critic(); }
  printf("Your must have looked it up!\n");
  return 0;
}

void critic() {
  printf("No luck, my friend, try again!\n");
  scanf("%d", &units);
}
