#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define BUFSIZE 4096
#define LENGTH 81

void append(FILE *src, FILE *dst);

char *str_gets(char *str, int n);

int main(int argc, char *argv[]) {
  FILE *fd, *fs;    // fd 指向目标文件，fs 指向来源文件
  int files = 0;    // 追加的文件数量
  char dst[LENGTH]; // 目标文件名称
  char src[LENGTH]; // 来源文件名称
  int ch;
  puts("Enter name of destination file (include full path): ");
  str_gets(dst, LENGTH);
  if ((fd = fopen(dst, "a+")) == NULL) {
    fprintf(stderr, "Can't open %s\n", dst);
    exit(EXIT_FAILURE);
  }
  if (setvbuf(fd, NULL, _IOFBF, BUFSIZE) != 0) {
    fputs("Can't create output buffer\n", stderr);
    exit(EXIT_FAILURE);
  }
  puts("Enter name of first source file (empty line to quit): ");
  while (str_gets(src, LENGTH) && src[0] != '\0') {
    if (strcmp(src, dst) == 0) {
      fputs("Can't append file to itself\n", stderr);
    } else if ((fs = fopen(src, "r")) == NULL) {
      fprintf(stderr, "Can't open %s\n", src);
    } else {
      if (setvbuf(fs, NULL, _IOFBF, BUFSIZE) != 0) {
        fputs("Can't create input buffer\n", stderr);
        continue;
      }
      append(fs, fd);
      if (ferror(fs) != 0) { fprintf(stderr, "Error in reading file %s\n", src); }
      if (ferror(fd) != 0) { fprintf(stderr, "Error in writing file %s\n", dst); }
      fclose(fs);
      files++;
      printf("File %s appended\n", src);
      puts("Enter name of next source file (empty line to quit): ");
    }
  }
  printf("Done appending. %d files appended\n", files);
  rewind(fd);
  printf("%s contents:\n", dst);
  while ((ch = getc(fd)) != EOF) { putchar(ch); }
  puts("Done displaying");
  fclose(fd);
  return 0;
}

void append(FILE *src, FILE *dst) {
  size_t bytes;
  static char temp[BUFSIZE]; // 只分配一次
  while ((bytes = fread(temp, sizeof(char), BUFSIZE, src)) > 0) { fwrite(temp, sizeof(char), bytes, dst); }
}

char *str_gets(char *str, int n) {
  char *res;
  int i = 0;
  res = fgets(str, n, stdin);
  if (res) {
    while (str[i] != '\n' && str[i] != '\0') { i++; }
    if (str[i] == '\n') {
      str[i] = '\0';
    } else {
      while (getchar() != '\n') { continue; }
    }
  }
  return res;
}
