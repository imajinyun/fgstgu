#include <stdio.h>

#define JUST_CHECKING
#define LIMIT 4
#define ABC

int main(int argc, char *argv[]) {
  int i;
  int total = 0;
  for (i = 0; i <= LIMIT; i++) {
    total += 2 * i * i + 1;
#ifdef JUST_CHECKING
    printf("i = %d, running total = %d\n", i, total);
#endif
  }
  printf("Total = %d\n\n", total);

#if defined(ICBC)
  printf("Industrial and commercial bank of china\n");
#elif defined(ABC)
  printf("Agricultural bank of china\n");
#elif defined(CCB)
  printf("China construction bank\n");
#else
  printf("Other bank\n");
#endif
  return 0;
}
