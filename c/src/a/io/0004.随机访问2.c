#include <stdio.h>
#include <stdlib.h>

#define ARRSIZE 1000

int main(int argc, char *argv[]) {
  double numbers[ARRSIZE];
  double value;
  const char *file = "/tmp/number.dat";
  int i;
  long pos;
  FILE *io;
  for (i = 0; i < ARRSIZE; i++) { numbers[i] = 100.0 * i + 1.0 / (i + 1); }
  if ((io = fopen(file, "wb")) == NULL) {
    fprintf(stderr, "Could not open %s for output\n", file);
    exit(EXIT_FAILURE);
  }
  fwrite(numbers, sizeof(double), ARRSIZE, io);
  fclose(io);
  if ((io = fopen(file, "rb")) == NULL) {
    fprintf(stderr, "Could not open %s for random access\n", file);
    exit(EXIT_FAILURE);
  }
  printf("Enter an index in the range 0-%d\n", ARRSIZE - 1);
  while (scanf("%d", &i) == 1 && i >= 0 && i < ARRSIZE) {
    pos = (long) i * sizeof(double); // 计算偏移量
    fseek(io, pos, SEEK_SET);        // 定位到此处
    fread(&value, sizeof(double), 1, io);
    printf("The value there is %f\n", value);
    printf("Enter next index (out of range to quit):\n");
  }
  fclose(io);
  puts("Done!");
  return 0;
}
