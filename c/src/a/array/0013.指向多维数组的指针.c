#include <stdio.h>

int main(int argc, char *argv[]) {
  int array[4][2] = {
    {2, 4},
    {6, 8},
    {1, 3},
    {5, 7},
  };
  int(*ptr)[2];
  ptr = array;

  // 可以用数组表示法或指针表示法来表示一个数组元素，既可以使用数组名，也可以使用指针名
  // ptr[m][n] = *(*(ptr + m) + n)
  // arr[m][n] = *(*(arr + m) + n)

  printf("   ptr = %p,    ptr + 1 = %p\n", ptr, ptr + 1);
  printf("ptr[0] = %p, ptr[0] + 1 = %p\n", ptr[0], ptr[0] + 1);
  printf("\n");
  printf("*ptr[0] = %d\n", ptr[0][0]);
  printf("*ptr[0] = %d\n", *ptr[0]);
  printf("  **ptr = %d\n", **ptr);
  printf("\n");
  printf("         pz[2][1] = %d\n", ptr[2][1]);
  printf("*(*(ptr + 2) + 1) = %d\n", *(*(ptr + 2) + 1));
  return 0;
}
