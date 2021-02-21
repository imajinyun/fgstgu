#include <stdbool.h>
#include <stdio.h>

// 线的样式
#define SOLID 0
#define DOTTED 1
#define DASHED 2

// 三原色
#define RED 1
#define GREEN 2
#define BLUE 4

// 混合色
#define BLACK 1
#define YELLOW (RED | GREEN)
#define MAGENTA (RED | BLUE)
#define CYAN (GREEN | BLUE)
#define WHITE (RED | GREEN | BLUE)

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

void show_settings(const struct box_props *box);

int main(int argc, char *argv[]) {
  struct box_props box = {true, YELLOW, true, GREEN, DASHED};
  printf("Original box settings:\n");
  show_settings(&box);
  box.opaque = false;
  box.fill_color = WHITE;
  box.border_color = MAGENTA;
  box.border_style = SOLID;
  printf("\nModified box settings:\n");
  show_settings(&box);
  return 0;
}

void show_settings(const struct box_props *bp) {
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
  }
}
