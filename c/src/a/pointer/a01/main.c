#include <stdio.h>

#define SIZE 4

int main(int argc, char *argv[]) {
  short dates[SIZE];
  short *ptrOfDate;
  short index;
  double bills[SIZE];
  double *ptrOfBill;

  // 把数组地址赋给指针
  ptrOfDate = dates;
  ptrOfBill = bills;
  printf("%s %13s %15s\n", "index", "short", "double");
  for (index = 0; index < SIZE; index++) {
    printf("pointers + %d: %10p %10p\n", index, ptrOfDate + index, ptrOfBill + index);
  }
  printf("\n");

  double array[SIZE] = {1.0, 2.0, 3.0, 4.0};

  // 相同的地址：array + 2 == &array[2]
  printf("array + 2: %p\n", array + 2);
  printf("&array[2]: %p\n", &array[2]);

  // 相同的取值：*(array + 2) == array[2]
  printf("*(array + 1): %.1f\n", *(array + 2));
  printf("array[2]: %.1f\n", array[2]);
  return 0;
}
