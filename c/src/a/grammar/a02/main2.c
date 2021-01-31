#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *itoa(int num, char *buffer, int base);
void *hexToBin(char *hex, char *bin);

// 浮点数结构体。
typedef struct {
  unsigned int nMant : 23; // 尾数部分
  unsigned int nExp : 8;   // 指数部分
  unsigned int nSign : 1;  // 符号位
} FP_SINGLE;

/**
 * 小数存储。
 */
int main(int argc, char *argv[]) {
  char s[10] = {0};
  char t[10] = {0};
  float f = 19.625;
  int x, y, z;
  char b1[65] = {0};
  char b2[65] = {0};
  char b3[65] = {0};
  FP_SINGLE *p = (FP_SINGLE *) &f;

  itoa(p->nSign, s, 2);
  printf("sign: %s\n", s);
  sprintf(t, "%x", p->nSign);
  hexToBin(t, b1);
  printf("sign: %s\n", b1);

  itoa(p->nExp, s, 2);
  printf("exp: %s\n", s);
  sprintf(t, "%x", p->nExp);
  hexToBin(t, b2);
  printf("exp: %s\n", b2);

  itoa(p->nMant, s, 2);
  printf("mant: %s\n", s);
  sprintf(t, "%x", p->nMant);
  hexToBin(t, b3);
  printf("mant: %s\n", b3);

  x = sprintf(s, "%d", p->nSign);
  y = sprintf(s, "%d", p->nExp);
  z = sprintf(s, "%d", p->nMant);
  printf("x = %d\n", x);
  printf("y = %d\n", y);
  printf("z = %d\n", z);
  printf("s = %s\n", s);

  printf("\n");
  char buffer[200], str[] = "computer", c = 'l';
  int i = 35, j = 0;
  float fp = 1.7320534f;

  // 格式化并打印各种数据到 buffer
  printf("%scharacter count = %d\n\n", buffer, j);
  j = sprintf(buffer, "String:    %s\n", str);
  j += sprintf(buffer + j, "Character: %c\n", c);
  j += sprintf(buffer + j, "Integer:   %d\n", i);
  j += sprintf(buffer + j, "Real:      %f\n", fp);
  printf("%s\ncharacter count = %d\n", buffer, j);
  return 0;
}

/**
 * 将整型值转换为字符串。
 *
 * @see [How do I use itoa() with GCC?](http://www.strudel.org.uk/itoa/)
 * @see [itoa Function in C](https://www.javatpoint.com/itoa-function-in-c)
 */
char *itoa(int num, char *buffer, int base) {
  int current = 0;
  if (num == 0) {
    buffer[current++] = '0';
    buffer[current] = '\0';
    return buffer;
  }
  int num_digits = 0;
  if (num < 0) {
    if (base == 10) {
      num_digits++;
      buffer[current] = '-';
      current++;
      num *= -1;
    }
    return NULL;
  }
  num_digits += (int) floor(log(num) / log(base)) + 1;
  while (current < num_digits) {
    int base_val = (int) pow(base, num_digits - 1 - current);
    int num_val = num / base_val;
    char value = num_val + '0';
    buffer[current] = value;
    current++;
    num -= base_val * num_val;
  }
  buffer[current] = '\0';
  return buffer;
}

void *hexToBin(char *hex, char *bin) {
  for (int i = 0; hex[i] != '\0'; i++) {
    switch (hex[i]) {
      case '0':
        strcat(bin, "0000");
        break;
      case '1':
        strcat(bin, "0001");
        break;
      case '2':
        strcat(bin, "0010");
        break;
      case '3':
        strcat(bin, "0011");
        break;
      case '4':
        strcat(bin, "0100");
        break;
      case '5':
        strcat(bin, "0101");
        break;
      case '6':
        strcat(bin, "0110");
        break;
      case '7':
        strcat(bin, "0111");
        break;
      case '8':
        strcat(bin, "1000");
        break;
      case '9':
        strcat(bin, "1001");
        break;
      case 'a':
      case 'A':
        strcat(bin, "1010");
        break;
      case 'b':
      case 'B':
        strcat(bin, "1011");
        break;
      case 'c':
      case 'C':
        strcat(bin, "1100");
        break;
      case 'd':
      case 'D':
        strcat(bin, "1101");
        break;
      case 'e':
      case 'E':
        strcat(bin, "1110");
        break;
      case 'f':
      case 'F':
        strcat(bin, "1111");
        break;
      default:
        printf("Invalid hexadecimal\n");
        break;
    }
  }
}
