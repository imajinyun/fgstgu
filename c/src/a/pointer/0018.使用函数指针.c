#include <ctype.h>
#include <stdio.h>
#include <string.h>

#define LEN 81

char *str_gets(char *str, int n);
char info();
void endline();
void show(void (*pfunc)(char *), char *str);
void ToUpper(char *);
void ToLower(char *);
void Transpose(char *);
void Dummy(char *);

int main(int argc, char *argv[]) {
  char line[LEN], copy[LEN];
  char choice;
  void (*pfunc)(char *); // 声明一个函数指针，被指向的函数接受 char * 类型的参数，无返回值
  puts("Enter a string (empty line to quit):");
  while (str_gets(line, LEN) != NULL && line[0] != '\0') {
    while ((choice = info()) != 'n') {
      switch (choice) {
        case 'u':
          pfunc = ToUpper;
          break;
        case 'l':
          pfunc = ToLower;
          break;
        case 't':
          pfunc = Transpose;
          break;
        case 'o':
          pfunc = Dummy;
          break;
      }
      strcpy(copy, line);
      show(pfunc, copy);
    }
    puts("Enter a string (empty line to quit):");
  }
  puts("Done!");
  return 0;
}

char info() {
  char ans;
  puts("Enter menu choice:");
  puts("u) uppercase  t) transpose case");
  puts("l) lowercase  o) original case");
  puts("n) next string");
  ans = tolower(getchar());
  endline(); // 清理输入行
  while (strchr("ulton", ans) == NULL) {
    puts("Please enter a [ulton]: ");
    ans = tolower(getchar());
    endline();
  }
  return ans;
}

void endline() {
  while (getchar() != '\n') { continue; }
}

void ToUpper(char *s) {
  while (*s) {
    *s = toupper(*s);
    s++;
  }
}

void ToLower(char *s) {
  while (*s) {
    *s = tolower(*s);
    s++;
  }
}

void Transpose(char *s) {
  while (*s) {
    if (islower(*s)) {
      *s = toupper(*s);
    } else if (isupper(*s)) {
      *s = tolower(*s);
    }
    s++;
  }
}

void Dummy(char *s) {}

void show(void (*pfunc)(char *), char *s) {
  (*pfunc)(s); // 把用户选定的函数作用于 s
  puts(s);
}

char *str_gets(char *str, int n) {
  char *res, *find;
  res = fgets(str, n, stdin);
  if (res) {
    if ((find = strchr(str, '\n')) != NULL) {
      *find = '\0';
    } else {
      while (getchar() != '\n') { continue; }
    }
  }
  return res;
}
