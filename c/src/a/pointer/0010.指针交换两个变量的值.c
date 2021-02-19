#include <stdio.h>

void swap(int *, int *);

int main(int argc, char *argv[]) {
  int a = 11, b = 99, tmp;
  int *pa = &a, *pb = &b;

  printf("初始的值：");
  printf("a = %d, b = %d, *pa = %d, *pb = %d\n", a, b, *pa, *pb);

  tmp = *pa;
  *pa = *pb;
  *pb = tmp;
  printf("简体互换：");
  printf("a = %d, b= %d, *pa = %d, *pb = %d\n", a, b, *pa, *pb);
  printf("\n");

  int c = 100, d = 200;
  printf("初始的值：");
  printf("c = %d, d = %d\n", c, d);
  printf("函数互换：");
  swap(&c, &d);
  printf("c = %d, d = %d\n", c, d);
  printf("\n");

  puts("假设有一个 int 类型的变量 a，pa 是指向它的指针，那么 *&a 和 &*pa 分别是什么意思：\n"
       "1. *&a：为 *(&a)，&a 表示取变量 a 的地址（等价于 pa），*(&a) 表示取这个地址上的数据（等价于 *pa），"
       "绕来绕去，又回到了原点，*&a 仍然等价于 a；\n"
       "2. &*pa：理解为 &(*pa)，*pa 表示取得 pa 指向的数据（等价于 a），&(*pa) 表示数据的地址（等价于 &a），"
       "所以 &*pa 等价于 pa；");
  return 0;
}

void swap(int *m, int *n) {
  int tmp;

  tmp = *m; // m 是 int *，*m 是 int
  *m = *n;
  *n = tmp;
}
