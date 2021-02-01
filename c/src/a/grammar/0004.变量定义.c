#include <stdio.h>

// 变量定义。
int main(int argc, char *argv[]) {
  char char1 = -128, char2 = 127;
  unsigned char char3 = 0, char4 = 255;
  signed char char5 = -128, char6 = 127;

  // %c 用来输出 char 类型，即一个字符，c 是 character 的简写
  printf("char1 = %c, char2 = %c\n", char1, char2);
  printf("char3 = %c, char4 = %c\n", char3, char4);
  printf("char5 = %c, char6 = %d\n", char5, char6);

  short int short1 = -32768, short2 = 32767;
  unsigned short short3 = 0, short4 = 65535;

  // %hd 用来输出 short int 类型，hd 是 short decimal 的简写
  printf("short1 = %hd, short2 = %hd\n", short1, short2);
  printf("short3 = %hd, short4 = %hd\n", short3, short4);

  int int1 = -2147483648, int2 = 2147483647;
  unsigned int int3 = 0, int4 = 4294967295;

  // %d 用来输出 int 类型，d 是 decimal 的简写
  printf("int1 = %d, int2 = %d\n", int1, int2);
  printf("int3 = %d, int4 = %d\n", int3, int4);

  long long1 = -2147483648, long2 = 2147483647;
  unsigned long long3 = 0, long4 = 4294967295;
  long long long5 = -2147483648, long6 = 2147483647;

  // %ld 用来输出 long int 类型，ld 是 long decimal 的简写
  printf("long1 = %ld, long2 = %ld\n", long1, long2);
  printf("long3 = %ld, long4 = %ld\n", long3, long4);
  printf("long5 = %lld, long6 = %lld\n", long5, long6);

  float float1 = 1.2E-38, float2 = 3.4E+38;

  // %f 用来输出 float 类型，f 是 float 的简写
  printf("float1 = %f, float2 = %f\n", float1, float2);

  double double1 = 2.3E-308, double2 = 1.7E+308;
  long double double3 = 3.4E-4932, double4 = 1.1E+4932;

  // %lf 用来输出 double 类型，lf 是 long float 的简写
  printf("double1 = %lf, double2 = %lf\n", double1, double2);
  printf("double3 = %Lf, double4 = %Lf\n", double3, double4);
  return 0;
}
