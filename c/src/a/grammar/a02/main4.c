#include <locale.h>
#include <stdio.h>
#include <wchar.h>

/**
 * 中文字符。
 */
int main(int argc, char *argv[]) {
  puts("putchar、printf 只能输出不加 L 前缀的窄字符，对加了 L 前缀的宽字符无能为力，\n"
       "我们必须使用 <wchar.h> 头文件中的宽字符输出函数，它们分别是 putwchar 和 wprintf\n");
  puts("putwchar 函数专门用来输出一个宽字符，它和 putchar 的用法类似\n"
       "wprintf 是通用的、格式化的宽字符输出函数，它除了可以输出单个宽字符，"
       "还可以输出宽字符串。宽字符对应的格式控制符为 %lc\n");

  wchar_t a = L'A';  // 英文字符（基本拉丁字符）
  wchar_t b = L'9';  // 英文数字（阿拉伯数字）
  wchar_t c = L'中'; // 中文汉字
  wchar_t d = L'国'; // 中文汉字
  wchar_t e = L'。'; // 中文标点
  wchar_t f = L'ヅ'; // 日文片假名
  wchar_t g = L'😄';  // 特殊符号
  wchar_t h = L'༄';  // 藏文

  setlocale(LC_ALL, "zh_CN");
  putwchar(a);
  putwchar(L' ');
  putwchar(b);
  putwchar(L' ');
  putwchar(c);
  putwchar(L' ');
  putwchar(d);
  putwchar(L' ');
  putwchar(e);
  putwchar(L' ');
  putwchar(f);
  putwchar(L' ');
  putwchar(g);
  putwchar(L' ');
  putwchar(h);
  putwchar(L'\n');

  wprintf(L"%lc %lc %lc %lc %lc %lc %lc %lc", a, b, c, d, e, f, g, h);

  puts("\n\n给字符串加上 L 前缀就变成了宽字符串，它包含的每个字符都是宽字符，一律采用 UTF-16 或者 UTF-32 编码。\n"
       "输出宽字符串可以使用 <wchar.h> 头文件中的 wprintf 函数，对应的格式控制符是 %ls\n");

  wchar_t str1[] = L"不加 L 前缀的窄字符串也可以处理中文";
  wchar_t *str2 = L"不加 L 前缀的窄字符串也可以处理中文";

  setlocale(LC_ALL, "zh_CN");
  wprintf(L"str1: %ls\nstr2: %ls\n", str1, str2);
  return 0;
}
