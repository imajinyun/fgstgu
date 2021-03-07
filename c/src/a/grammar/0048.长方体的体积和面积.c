#include <stdio.h>

float f1, f2, f3;

float cuboidVolume(float length, float width, float height);

int main(int argc, char *argv[]) {
  float v, length, width, height;
  printf("请输入长方体的长宽高：");
  scanf("%f %f %f", &length, &width, &height);

  v = cuboidVolume(length, width, height);
  printf("体积 = %.2f\n", v);
  printf("顶部（或者底部）面积（长*宽）：%.2f\n", f1);
  printf("正面（或者背面）面积（宽*高）：%.2f\n", f2);
  printf("左侧（右侧）面积（长*高）：%.2f\n", f3);
  return 0;
}

float cuboidVolume(float length, float width, float height) {
  float v;
  v = length * width * height; // 体积
  f1 = length * width;         // 顶部（或者底部）面积
  f2 = width * height;         // 正面（或者背面）面积
  f3 = length * height;        // 左侧（右侧）面积

  return v;
}
