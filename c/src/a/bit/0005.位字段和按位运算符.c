#include <limits.h>
#include <stdbool.h>
#include <stdio.h>

// 边框线样式
#define SOLID 0
#define DOTTED 1
#define DASHED 2

// 三原色
#define RED 1
#define GREEN 2
#define BLUE 4

// 混合色
#define BLACK 0
#define YELLOW (RED | GREEN)
#define MAGENTA (RED | BLUE)
#define CYAN (GREEN | BLUE)
#define WHITE (RED | GREEN | BLUE)

// 按位方法中用到的符号常量
#define OPAQUE 0x1
#define FILL_RED 0x2
#define FILL_GREEN 0x4
#define FILL_BLUE 0x8
#define FILL_MASK 0xE
#define BORDER 0x100
#define BORDER_RED 0x200
#define BORDER_GREEN 0x400
#define BORDER_BLUE 0x800
#define BORDER_MASK 0xE00
#define B_SOLID 0
#define B_DOTTED 0x1000
#define B_DASHED 0x2000
#define STYLE_MASK 0x3000

const char *colors[8] = {"black", "red", "green", "yellow", "blue", "magenta", "cyan", "white"};

struct box_props {
  bool opaque : 1;
  unsigned int fill_color : 3;
  unsigned int : 4;
  bool show_border : 1;
  unsigned int border_color : 3;
  unsigned int border_style : 2;
  unsigned int : 2;
};

// 把数据看作结构或 unsigned short 类型的变量
union views {
  struct box_props st_view;
  unsigned short us_view;
};

void show_st_view(const struct box_props *bp);
void show_us_view(unsigned short);
char *itobs(int n, char *bs);

int main(int argc, char *argv[]) {
  union views box = {{true, YELLOW, true, GREEN, DASHED}};
  char bs[8 * sizeof(unsigned int) + 1];
  printf("Original box show:\n");
  show_st_view(&box.st_view);
  printf("\nBox show using unsigned int view:\n");
  show_us_view(box.us_view);
  printf("bits are %s\n", itobs(box.us_view, bs));
  box.us_view &= ~FILL_MASK;               // 把表示填充色的位清 0
  box.us_view |= (FILL_BLUE | FILL_GREEN); // 重置填充色
  box.us_view ^= OPAQUE;                   // 切换是否透明的位
  box.us_view |= BORDER_RED;               // 错误的方法
  box.us_view &= ~STYLE_MASK;              // 把样式位清 0
  box.us_view |= B_DOTTED;                 // 把样式设置为点
  printf("\nModified box show:\n");
  show_st_view(&box.st_view);
  printf("\nBox show using unsigned int view:\n");
  show_us_view(box.us_view);
  printf("bits are %s.\n", itobs(box.us_view, bs));
  return 0;
}

void show_st_view(const struct box_props *bp) {
  printf("Box is %s.\n", bp->opaque == true ? "opaque" : "transparent");
  printf("The fill color is %s.\n", colors[bp->fill_color]);
  printf("Border %s.\n", bp->show_border == true ? "shown" : "not shown");
  printf("The border color is %s.\n", colors[bp->border_color]);
  printf("The border style is ");
  switch (bp->border_style) {
    case SOLID:
      printf("solid.\n");
      break;
    case DOTTED:
      printf("dotted.\n");
      break;
    case DASHED:
      printf("dashed.\n");
      break;
    default:
      printf("unknown type.\n");
      break;
  }
}

void show_us_view(unsigned short ut) {
  printf("box is %s\n", (ut & OPAQUE) == OPAQUE ? "opaque" : "transparent");
  printf("The fill color is %s\n", colors[(ut >> 1) & 07]);
  printf("Border %s\n", (ut & BORDER) == BORDER ? "shown" : "not shown");
  printf("The border style is ");
  switch (ut & STYLE_MASK) {
    case B_SOLID:
      printf("solid.\n");
      break;
    case B_DOTTED:
      printf("dotted.\n");
      break;
    case B_DASHED:
      printf("dashed.\n");
      break;
    default:
      printf("unknown type.\n");
      break;
  }
  printf("The border color is %s.\n", colors[(ut >> 9) & 07]);
}

char *itobs(int n, char *bs) {
  int i;
  const static int size = CHAR_BIT * sizeof(int);
  for (i = size - 1; i >= 0; i--, n >>= 1) { bs[i] = (01 & n) + '0'; }
  bs[size] = '\0';
  return bs;
}
