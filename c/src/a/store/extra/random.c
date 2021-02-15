#include <stdio.h>

static unsigned long int next = 1; // 种子

void srandom(unsigned int seed) { next = seed; }

/**
 * 静态变量 next 的初始值是 1，其值在每次调用 random() 函数时都会被修改（通过魔术公式）。
 * 该函数是用于返回一个 0～32767 之间的值。
 *
 * 注意：next 是具有内部链接的静态变量（并非无链接）。这是为了方便扩展，供同一个文件中的其他函数共享。
 */
unsigned int random() {
  // 生成伪随机数的魔术公式
  next = next * 1103515245 + 12345;
  return (unsigned int) (next / 65535) % 32768;
}
