#include <stdio.h>

extern void srandom(unsigned int x);
extern unsigned int random();

int main(int argc, char *argv[]) {
  unsigned int seed;
  printf("Please enter your choice for seed: ");
  while (scanf("%u", &seed) == 1) {
    srandom(seed);
    for (int i = 0; i < 5; i++) { printf("random = %d\n", random()); }
    printf("Please enter next seed (q to quit): ");
  }
  printf("Done!\n");
  return 0;
}
