#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define LENGTH 40

/**
 * 压缩文件内容（将文件压缩为原来的三分之一）。
 *
 * -> vim /tmp/test.txt
 *    hello world, hello world, this is a test case.
 * -> /path/to/a.io.0002 /tmp/test.txt hello
 */
int main(int argc, char *argv[]) {
  FILE *in, *out;
  int ch;
  char name[LENGTH];
  int count = 0;
  if (argc < 2) {
    fprintf(stderr, "Usage: %s filename\n", argv[0]);
    exit(EXIT_FAILURE);
  }
  if ((in = fopen(argv[1], "r")) == NULL) {
    fprintf(stderr, "I couldn't open the file \"%s\"\n", argv[1]);
    exit(EXIT_FAILURE);
  }
  strncpy(name, argv[1], LENGTH - 5);
  name[LENGTH - 5] = '\0';
  fprintf(stdout, "strncpy after: %s\n", name);
  strcat(name, ".zip");
  fprintf(stdout, "strcat after: %s\n", name);

  if ((out = fopen(name, "w")) == NULL) {
    fprintf(stderr, "Can't create output file\n");
    exit(EXIT_FAILURE);
  }
  while ((ch = getc(in)) != EOF) {
    if (count++ % 3 == 0) { putc(ch, out); }
  }
  if (fclose(in) != 0 || fclose(out) != 0) {
    fprintf(stderr, "Error in closing files\n");
  }
  return 0;
}
