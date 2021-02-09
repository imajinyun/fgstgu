#include <stdio.h>

#define LENGTH 40
#define LIMIT 5

/**
 * 字符串数组：
 *
 * 1. 如果要用数组表示一系列待显示的字符串，请使用指针数组，因为它比二维字符数组的效率高；
 * 2. 如果要改变字符串或为字符串输入预留空间，不要使用指向字符串字面量的指针；
 */
int main(int argc, char *argv[]) {
  // a 数组是一个内含 5 个指针的数组，在系统中共占用 40 字节。
  // a 中的指针指向初始化时所用的字符串字面量的位置，这些字符串字面量被存储在静态内存中。
  const char *a[LIMIT] = {"Adding numbers swiftly", "Multiplying accurately", "Stashing data",
                          "Following instructions to the letter", "Understanding the C language"};

  // b 数组是一个内含 5 个数组的数组，每个数组内含 40 个 char 类型的值，共占用 200 字节。
  // b 数组中的数组则存储着字符串字面量的副本，所以每个字符串都被存储了两次。
  char b[LIMIT][LENGTH] = {
    "Walking in a straight line", "Sleeping", "Watching television", "Mailing letters", "Reading email",
  };

  // 为字符串数组分配内存的使用率较低。b 中的每个元素的大小必须相同，而且必须是能存储最长字符串的大小。
  puts("Let's compare a vs b:");
  printf("%-36s %-25s\n", "A", "B");
  for (int i = 0; i < LIMIT; i++) {
    printf("%-36s %-25s\n", a[i], b[i]);
    printf("sizeof a: %zd, sizeof b: %zd\n\n", sizeof(a), sizeof(b));
  }
  return 0;
}
