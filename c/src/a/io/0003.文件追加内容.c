#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX 41

int main(int argc, char *argv[]) {
  FILE *fp;
  char words[MAX];
  if ((fp = fopen("/tmp/test.txt", "a+")) == NULL) {
    fprintf(stdout, "Can't open \"/tmp/test.txt\" file\n");
    exit(EXIT_FAILURE);
  }
  puts("Enter words to add to the file; press the #");
  puts("key at the beginning of a line to terminate");
  while ((fscanf(stdin, "%40s", words) == 1) && (words[0] != '#')) { fprintf(fp, "%s\n", words); }
  puts("File contents:");
  rewind(fp); // 返回到文件开始处
  while (fscanf(fp, "%s", words) == 1) { puts(words); }
  puts("Done!\n");
  if (fclose(fp) != 0) { fprintf(stderr, "Error closing file\n"); }
  return 0;
}
