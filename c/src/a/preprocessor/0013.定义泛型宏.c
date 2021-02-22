#include <math.h>
#include <stdio.h>

#define RAD_TO_DEC (180 / (4 * atanl(1)))

// 泛型平方根函数
#define SQRT(x) _Generic((x), long double : sqrtl, default : sqrt, float : sqrtf)(x)

// 泛型正弦函数，角度的单位为度
#define SIN(x)                                                                                                         \
  _Generic((x), long double : sinl(x) / RAD_TO_DEC, default : sin((x) / RAD_TO_DEC), float : sinf((x) / RAD_TO_DEC))

int main(int argc, char *argv[]) {
  float x1 = 45.0f;
  double x2 = 45.0;
  long double x3 = 45.0L;
  long double y1 = SQRT(x1);
  long double y2 = SQRT(x2);
  long double y3 = SQRT(x3);
  printf("%.17Lf\n", y1); // 匹配 float
  printf("%.17Lf\n", y2); // 匹配 default
  printf("%.17Lf\n", y3); // 匹配 long double

  int i = 45;
  y2 = SQRT(i);
  printf("%.17Lf\n", y2); // 匹配 default
  y3 = SIN(x3);
  printf("%.17Lf\n", y3); // 匹配 long double
}
